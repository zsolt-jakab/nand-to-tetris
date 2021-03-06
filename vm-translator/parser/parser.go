package parser

import (
	"fmt"
	"github.com/zsolt-jakab/nand-to-tetris/vm-translator/parser/command/arithlog"
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
func Translate(fileName string, codeLines []string, codeLineIndexes []int) []string {
	var translatedFileLines []string
	var translatedLineOfCommand []string
	//translatedLineOfCommand := make([]string, 0)
	for index, codeLine := range codeLines {
		var err error
		comm := strings.Fields(codeLine)
		if len(comm) == 1 {
			translatedLineOfCommand, err = arithlog.Translate(comm[0])
		} else if len(comm) == 3 {
			translatedLineOfCommand, err = memaccess.Translate(memaccess.NewBase(comm[0], comm[1], comm[2]), fileName)
		}
		if err != nil {
			panic(fmt.Sprintf("Error: [%v] in code: [%s] in line: [%d]", err, codeLine, codeLineIndexes[index]))
		}
		translatedFileLines = append(translatedFileLines, translatedLineOfCommand...)
	}
	return translatedFileLines
}
