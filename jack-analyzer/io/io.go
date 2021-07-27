package io

import (
	"regexp"
	"strings"
)

const (
	newLine = "\n"
	emptyString = ""
)

/*
FileAccessor is an interface for reading code lines of source file and creating new files
*/
type FileAccessor interface {
	ReadCodeAsLines(name string) ([]string, []int)
	ReadCode(name string) string
	CreateFile(name string, lines []string)
}

/*
DefaultFileAccessor is the base implementation of FileAccessor
*/
type DefaultFileAccessor struct {
	FileReader
	FileWriter
}

/*
ReadCodeAsLines reads the code lines from a source file into a string array.
It will ignore all of the empty lines and the comments.
It will give back the line numbers in the files in an int array, what will help later the parser to locate the error,
if a there is an invalid code which can not be translated to binary.
*/
func (sr *DefaultFileAccessor) ReadCodeAsLines(name string) ([]string, []int) {
	bytes, err := sr.Read(name)
	panicIfError(err)
	codeSrc := stripComments(string(bytes))
	lines := strings.Split(codeSrc, newLine)

	return scanRawLines(lines)
}

/*
ReadCode reads the code lines from a source file into a string array.
It will ignore all of the empty lines and the comments.
It will give back the line numbers in the files in an int array, what will help later the parser to locate the error,
if a there is an invalid code which can not be translated to binary.
*/
func (sr *DefaultFileAccessor) ReadCode(name string) string {
	codeAsLines, _ := sr.ReadCodeAsLines(name)

	return join(codeAsLines)
}

/*
CreateFile creates a file with the given name and extension and lines as content
*/
func (sr *DefaultFileAccessor) CreateFile(name string, lines []string) {
	joinedLines := join(lines)
	err := sr.Write(name, []byte(joinedLines))

	panicIfError(err)
}

func scanRawLines(lines []string) ([]string, []int) {
	lineCount := 0
	var codeLines []string
	var codeLineNumbers []int
	for _, line := range lines {
		lineCount++
		instruction := getInstructionPart(line)
		if instruction != "" {
			codeLines = append(codeLines, instruction)
			codeLineNumbers = append(codeLineNumbers, lineCount)
		}
	}

	return codeLines, codeLineNumbers
}

func getInstructionPart(line string) string {
	return strings.TrimSpace(line)
}

func stripComments(src string) string {
	multiLineComments := regexp.MustCompile("(?s)/\\*.*?\\*/")
	singleLineComments := regexp.MustCompile("//.*")
	return singleLineComments.ReplaceAllString(multiLineComments.ReplaceAllString(src, emptyString), emptyString)
}

func join(lines []string) string {
	return strings.Join(lines, newLine) + newLine
}

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}
