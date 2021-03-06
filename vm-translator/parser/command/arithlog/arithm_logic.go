package arithlog

import (
	"fmt"
	"strconv"
	"strings"
)

var nextId = idSequence()

var operationMap = map[string][]string{
	//arithmetic commands
	"add": {"@SP", "AM=M-1", "D=M", "A=A-1", "M=D+M"},
	"sub": {"@SP", "AM=M-1", "D=M", "A=A-1", "M=M-D"},
	"neg": {"@SP", "A=M-1", "M=-M"},

	//logical commands
	"and": {"@SP", "AM=M-1", "D=M", "A=A-1", "M=D&M"},
	"or":  {"@SP", "AM=M-1", "D=M", "A=A-1", "M=D|M"},
	"not": {"@SP", "A=M-1", "M=!M"},
}

var comparisonMap = map[string]func() []string{
	"eq": comparison("JEQ"),
	"gt": comparison("JGT"),
	"lt": comparison("JLT"),
}

func Translate(command string) ([]string, error) {
	return getTranslation(command)
}

func getTranslation(command string) ([]string, error) {
	if translation, isPresent := operationMap[command]; isPresent {
		return translation, nil
	} else if translation, isPresent := comparisonMap[command]; isPresent {
		return translation(), nil
	}
	return nil, fmt.Errorf("This is not a valid arithmetic or logical command : [%s] ", command)
}

func comparison(branchJumpCondition string) func() []string {
	//we need this id to make sure we will create new labels every time so we don't jump back to some previous label
	return func() []string {
		id := nextId()
		return []string{
			//D = x - y
			"@SP",
			"AM=M-1",
			"D=M",
			"A=A-1",
			"D=M-D",

			//if jump condition true for D(eq, lt, gt 0), jump
			"@" + strings.ToLower(branchJumpCondition) + "_true_case" + id,
			"D;" + branchJumpCondition,

			//false case
			"@SP",
			"A=M-1",
			"M=0",
			"@" + strings.ToLower(branchJumpCondition) + "_end" + id,
			"0;JMP",

			//true case
			"(" + strings.ToLower(branchJumpCondition) + "_true_case" + id + ")",
			"@SP",
			"A=M-1",
			"M=-1",

			//end of if
			"(" + strings.ToLower(branchJumpCondition) + "_end" + id + ")",
		}
	}
}

func idSequence() func() string {
	id := 0
	return func() string {
		id++
		return strconv.Itoa(id)
	}
}
