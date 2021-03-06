package io

import (
	"strings"
)

const (
	asmFileExtension = ".asm"
	vmFileExtension  = ".vm"
	newLine          = "\n"
)

/*
FileAccessor is an interface for reading code lines of a source file and creating new files with translated content
*/
type FileAccessor interface {
	ReadSourceLines(name string) ([]string, []int)
	CreateTargetFile(name string, lines []string)
}

/*
VMTranslatorFileAccessor is the base implementation of FileAccessor
*/
type VMTranslatorFileAccessor struct {
	FileReader
	FileWriter
}

/*
ReadSourceLines reads the code lines from a .vm file into a string array.
It will ignore all of the empty lines and the comments.
It will give back the line numbers in the files in an int array, what will help the parser to locate the error,
if a there is an invalid code which can not be translated.
*/
func (sr *VMTranslatorFileAccessor) ReadSourceLines(name string) ([]string, []int) {
	bytes, err := sr.Read(name + vmFileExtension)
	panicIfError(err)
	lines := strings.Split(string(bytes), newLine)

	return scanRawLines(lines)
}

/*
CreateTargetFile creates a file with the given lines and .vm extension
*/
func (sr *VMTranslatorFileAccessor) CreateTargetFile(name string, lines []string) {
	joinedLines := join(lines)
	err := sr.Write(name+asmFileExtension, []byte(joinedLines))

	panicIfError(err)
}

func scanRawLines(lines []string) ([]string, []int) {
	lineCount := 0
	var codeLines []string
	var codeLineIndexes []int
	for _, line := range lines {
		lineCount++
		instruction := getInstructionPart(line)
		if instruction != "" {
			codeLines = append(codeLines, instruction)
			codeLineIndexes = append(codeLineIndexes, lineCount)
		}
	}

	return codeLines, codeLineIndexes
}

func getInstructionPart(line string) string {
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
