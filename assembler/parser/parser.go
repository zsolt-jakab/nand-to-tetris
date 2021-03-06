package parser

import (
	"fmt"
	"strings"

	"github.com/zsolt-jakab/nand-to-tetris/assembler/parser/instruction"
)

const (
	emptyJump = "null"
	emptyDest = "null"
)

/*
Instruction interface has one function, the Binary() function
It should return the binary representation of an instruction
*/
type Instruction interface {
	Binary() string
}

/*
Translate translates lines of hack assembly code into binary code lines
It has a 2nd parameter, the count of that line in the original source file, what should
help in case of some error happens in the translation
*/
func Translate(codeLines []string, codeLineIndexes []int) []string {
	symbolTable := NewSymbolTable(codeLines)
	var binaryLines []string
	for indx, code := range codeLines {
		if !isLabel(code) {
			instruction, err := getInstruction(code, symbolTable)
			if err != nil {
				panic(fmt.Sprintf("Error: [%v] in code: [%s] in line: [%d]", err, code, codeLineIndexes[indx]))
			}
			binaryLines = append(binaryLines, instruction.Binary())
		}
	}
	return binaryLines
}

func getInstruction(inst string, symbolTable *SymbolTable) (Instruction, error) {
	if isAInstruction(inst) {
		return getAInstruction(inst, symbolTable)
	}
	return getCInstruction(inst)
}

func getAInstruction(inst string, symbolTable *SymbolTable) (Instruction, error) {
	addressText := trimAInstruction(inst)
	address, err := symbolTable.GetAddress(addressText)
	if err != nil {
		return nil, err
	}
	return instruction.NewA(address)
}

func getCInstruction(inst string) (Instruction, error) {
	dest := getDestination(inst)
	jump := getJump(inst)
	comp := getComputation(inst)

	return instruction.NewC(dest, comp, jump)
}

func getComputation(instruction string) string {
	beginOfComp := strings.Index(instruction, "=") + 1

	endOfComp := len(instruction)
	if strings.Contains(instruction, ";") {
		endOfComp = strings.Index(instruction, ";")
	}

	comp := string(instruction[beginOfComp:endOfComp])

	return comp
}

func getJump(instruction string) string {
	var jump string
	if strings.Contains(instruction, ";") {
		beginOfJump := strings.Index(instruction, ";") + 1
		jump = string(instruction[beginOfJump:])
	}
	return jump
}

func getDestination(instruction string) string {
	var dest string
	if strings.Contains(instruction, "=") {
		endOfDest := strings.Index(instruction, "=")
		dest = string(instruction[0:endOfDest])
	}
	return dest
}
