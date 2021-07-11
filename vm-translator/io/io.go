package io

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	asmFileExtension = ".asm"
	vmFileExtension  = ".vm"
	vmFilePattern    = "*" + vmFileExtension
	newLine          = "\n"
)

/*
CodeLine is the representation of a code line, with the Content of a code line, the file name of that line and
the Index of that line.
*/
type CodeLine struct {
	FileName string
	Content  string
	Index    int
}

/*
FileAccessor is an interface for reading code lines of a source file and creating new files with translated Content
*/
type FileAccessor interface {
	ReadSourceLines(name string) []CodeLine
	CreateTargetFile(name string, lines []string)
}

/*
VMTranslatorFileAccessor is the base implementation of FileAccessor
*/
type VMTranslatorFileAccessor struct {
	FileReader
	FileWriter
	FileAccess
}

/*
ReadSourceLines reads the code lines from a .vm file(s) into a string array.
It will ignore all of the empty lines and the comments.
It will give back the line numbers in the files in an int array, what will help the parser to locate the error,
if a there is an invalid code which can not be translated.
*/
func (sr *VMTranslatorFileAccessor) ReadSourceLines(filePath string) []CodeLine {
	fileInfo, _ := sr.Stat(filePath)
	var codeLines []CodeLine
	if fileInfo.IsDir() {
		var vmFilePaths, err = findPaths(filePath, vmFilePattern)
		panicIfError(err)
		fmt.Println("filePaths: " + join(vmFilePaths))
		for _, vmFilePath := range vmFilePaths {
			fmt.Println("filePath: " + vmFilePath)
			bytes, err := sr.Read(vmFilePath)
			panicIfError(err)
			lines := strings.Split(string(bytes), newLine)
			fmt.Println("lines: " + join(lines))
			codeLines = append(codeLines, scanRawLines(vmFilePath, lines)...)
		}
	} else {
		bytes, err := sr.Read(filePath)
		panicIfError(err)
		lines := strings.Split(string(bytes), newLine)
		codeLines = append(codeLines, scanRawLines(filePath, lines)...)
	}
	return codeLines
}

/*
CreateTargetFile creates a file with the given lines and .asm extension
*/
func (sr *VMTranslatorFileAccessor) CreateTargetFile(filePath string, lines []string) {
	var joinedLines = join(lines)
	var filePathToWrite string
	fileInfo, _ := sr.Stat(filePath)
	if fileInfo.IsDir() {
		filePathToWrite = path.Join(filePath, path.Base(filePath)+asmFileExtension)
	} else {
		filePathToWrite = strings.TrimRight(filePath, vmFileExtension) + asmFileExtension
	}

	err := sr.Write(filePathToWrite, []byte(joinedLines))

	panicIfError(err)
}

func scanRawLines(vmFilePath string, rawLines []string) []CodeLine {
	fileName := strings.TrimRight(filepath.Base(vmFilePath), vmFileExtension)
	lineCount := 0
	var codeLines []CodeLine
	for _, rawLine := range rawLines {
		lineCount++
		codePart := getCodePart(rawLine)
		if codePart != "" {
			codeLines = append(codeLines, CodeLine{FileName: fileName, Content: codePart, Index: lineCount})
		}
	}
	return codeLines
}

func getCodePart(line string) string {
	instructionPart := stripComment(line)
	return strings.TrimSpace(instructionPart)
}

func stripComment(line string) string {
	return strings.Split(line, startOfComment)[0]
}

func join(lines []string) string {
	return strings.Join(lines, newLine) + newLine
}

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func findPaths(root, pattern string) ([]string, error) {
	fmt.Println("root : " + root)
	fmt.Println("pattern : " + pattern)
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fmt.Println("walk : " + path)
		if err != nil {
			fmt.Println("err : " + err.Error())
			return err
		}
		if info.IsDir() {
			fmt.Println("it is a Dir : ")
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			fmt.Println("No match")
			return err
		} else if matched {
			matches = append(matches, path)
			fmt.Println("matches ", matches)
		}
		fmt.Println("nothing ")
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}
