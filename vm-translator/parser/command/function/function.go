package function

import (
	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/parser/command/util"
	"strconv"
)

var nextId = util.IdSequence()

func CreateFunction(functionName string, nVarsText string) []string {
	var nVars, _ = strconv.Atoi(nVarsText)
	var codeLines []string
	var pushZeroToStackCommands = pushZero()
	codeLines = append(codeLines, "("+functionName+")")
	for i := 0; i < nVars; i++ {
		codeLines = append(codeLines, pushZeroToStackCommands...)
	}
	return codeLines
}

func CallFunction(functionName string, nArgsText string) []string {
	var nArgs, _ = strconv.Atoi(nArgsText)
	var retAddrLabel = functionName + "$ret." + nextId()

	var codeLines []string
	// push retAddrLabel (Using a translator-generated label)
	codeLines = append(codeLines, "@"+retAddrLabel, "D=A", "@SP", "A=M", "M=D", "@SP", "M=M+1")

	// push LCL (Saves LCL of the caller)
	codeLines = append(codeLines, "@LCL", "D=M", "@SP", "A=M", "M=D", "@SP", "M=M+1")

	// push ARG (Saves ARG of the caller)
	codeLines = append(codeLines, "@ARG", "D=M", "@SP", "A=M", "M=D", "@SP", "M=M+1")

	// push THIS (Saves THIS of the caller)
	codeLines = append(codeLines, "@THIS", "D=M", "@SP", "A=M", "M=D", "@SP", "M=M+1")

	// push THAT (Saves THAT of the caller)
	codeLines = append(codeLines, "@THAT", "D=M", "@SP", "A=M", "M=D", "@SP", "M=M+1")

	// ARG = SP-5-nArgs (Repositions ARG)
	var diff = strconv.Itoa(5 + nArgs)
	codeLines = append(codeLines, "@SP", "D=M", "@"+diff, "D=D-A", "@ARG", "M=D")

	// LCL = SP (Repositions LCL)
	codeLines = append(codeLines, "@SP", "D=M", "@LCL", "M=D")

	// goto functionName (Transfers control to the called function)
	codeLines = append(codeLines, "@"+functionName, "0;JMP")

	// (retAddrLabel) (the same translator-generated label)
	codeLines = append(codeLines, "("+retAddrLabel+")")

	return codeLines
}

func ReturnCommand() []string {
	var codeLines []string
	// endFrame = LCL (endFrame is a temporary variable)
	codeLines = append(codeLines, "@LCL", "D=M", "@endFrame", "M=D")

	//retAddr = *(endFrame – 5) (gets the return address)
	codeLines = append(codeLines, "@endFrame", "D=M", "@5", "D=D-A", "A=D", "D=M", "@retAddr", "M=D")

	// *ARG = pop() (repositions the return value for the caller)
	codeLines = append(codeLines, "@SP", "AM=M-1", "D=M", "@ARG", "A=M", "M=D")

	// SP = ARG + 1 (repositions SP of the caller)
	codeLines = append(codeLines, "@ARG", "D=M+1", "@SP", "M=D")

	//THAT = *(endFrame – 1) (restores THAT of the caller)
	codeLines = append(codeLines, "@endFrame", "D=M", "@1", "D=D-A", "A=D", "D=M", "@THAT", "M=D")

	//THIS = *(endFrame – 2) (restores THIS of the caller)
	codeLines = append(codeLines, "@endFrame", "D=M", "@2", "D=D-A", "A=D", "D=M", "@THIS", "M=D")

	//ARG = *(endFrame – 3) (restores ARG of the caller)
	codeLines = append(codeLines, "@endFrame", "D=M", "@3", "D=D-A", "A=D", "D=M", "@ARG", "M=D")

	//LCL = *(endFrame – 4) (restores LCL of the caller)
	codeLines = append(codeLines, "@endFrame", "D=M", "@4", "D=D-A", "A=D", "D=M", "@LCL", "M=D")

	//goto retAddr (goes to the caller’s return address)
	codeLines = append(codeLines, "@retAddr", "A=M", "0;JMP")

	return codeLines
}

func pushZero() []string {
	return []string{
		// store the value of 0 in D
		"@0",
		"D=A",

		// store D(0) to *SP
		"@SP",
		"A=M",
		"M=D",

		// SP++
		"@SP",
		"M=M+1",
	}
}
