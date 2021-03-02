package io

import (
	"strings"
)

const (
	hackFileExtension = ".hack"
	newLine           = "\n"
)

/*
FileAccessor is an interface for reading code lines of the hack assembly file and creating new hack binary files
*/
type FileAccessor interface {
	ReadCodeLines(name string) ([]string, []int)
	CreateHackFile(name string, lines []string)
}

/*
DefaultFileAccessor is the base implementation of FileAccessor
*/
type DefaultFileAccessor struct {
	FileReader
	FileWriter
}

/*
ReadCodeLines reads the code lines from a Hack assembly(asm) file into a string array.
It will ignore all of the empty lines and the comments.
It will give back the line numbers in the files in an int array, what will help the parser to locate the error,
if a there is an invalid assembly code which can not be translated to binary.
*/
func (sr *DefaultFileAccessor) ReadCodeLines(name string) ([]string, []int) {
	bytes, err := sr.Read(name + ".asm")
	panicIfError(err)
	lines := strings.Split(string(bytes), newLine)

	return scanRawLines(lines)
}

/*
CreateHackFile creates a file with the given lines and .hack extension
Lines should contain binary code what can run in the hack computer.
*/
func (sr *DefaultFileAccessor) CreateHackFile(name string, lines []string) {
	joinedLines := join(lines)
	err := sr.Write(name+hackFileExtension, []byte(joinedLines))

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
