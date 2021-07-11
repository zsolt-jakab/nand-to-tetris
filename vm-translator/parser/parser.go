package parser

import (
	"fmt"
	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/io"
	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/parser/command/arithlog"
	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/parser/command/branching"
	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/parser/command/function"
	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/parser/command/memaccess"
	"strings"
)

/*
Command interface has one function, the Translate() function
It should return the assembly representation of the command
*/
type Command interface {
	Translate() ([]string, error)
}

/*
Translate translates lines of vm code into hack assembly code lines
It has a 2nd parameter, the count of that line in the original source file, what should
help in case of some error happens in the translation
*/
func Translate(codeLines []io.CodeLine, bootStrap bool) []string {
	var currentFunction string
	var translatedFileLines []string
	var translatedLineOfCommand []string
	if bootStrap {
		translatedFileLines = append(translatedFileLines, "@256", "D=A", "@SP", "M=D")
		translatedFileLines = append(translatedFileLines, function.CallFunction("Sys.init", "0")...)
	}
	for _, codeLine := range codeLines {
		var err error
		comm := strings.Fields(codeLine.Content)
		translatedFileLines = append(translatedFileLines, "// "+codeLine.Content)
		if comm[0] == "function" {
			currentFunction = comm[1] + "$"
			translatedLineOfCommand = function.CreateFunction(comm[1], comm[2])
		} else if comm[0] == "call" {
			translatedLineOfCommand = function.CallFunction(comm[1], comm[2])
		} else if comm[0] == "return" {
			translatedLineOfCommand = function.ReturnCommand()
		} else if len(comm) == 1 {
			translatedLineOfCommand, err = arithlog.Translate(comm[0])
		} else if len(comm) == 2 {
			translatedLineOfCommand, err = branching.Translate(comm[0], currentFunction+comm[1])
		} else if len(comm) == 3 {
			translatedLineOfCommand, err = memaccess.Translate(memaccess.NewBase(comm[0], comm[1], comm[2]), codeLine.FileName)
		}
		if err != nil {
			panic(fmt.Sprintf("Error: [%v] in file: [%s] in code: [%s] in line: [%d]", err, codeLine.FileName, codeLine.Content, codeLine.Index))
		}
		translatedFileLines = append(translatedFileLines, translatedLineOfCommand...)
	}
	return translatedFileLines
}
