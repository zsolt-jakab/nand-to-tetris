package parser

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	firstLineIndex   = 0
	firstReeRegister = 16
)

/*
SymbolTable is wrapper type around a string int map of symbols
There are 3 kind of symbols in a SymbolTable
	- 23 pre defined symbols
	- labels or the source code
	- variables of the source code
*/
type SymbolTable struct {
	symbols map[string]int
}

/*
NewSymbolTable creates a new SymbolTable with all of the
pre defined symbols, labels and variables from the lines of the source code.
*/
func NewSymbolTable(lines []string) *SymbolTable {
	symbolTable := initializeSymbolTable()
	symbolTable.addLabels(lines)
	symbolTable.addVariables(lines)

	return symbolTable
}

/*
GetAddress returns the address for a symbol.
If the symbol was already an address, it just converts it
If the symbol is not a number, it will try to get it from the map.
Raise an Error if can not find the symbol.
*/
func (symbolTable *SymbolTable) GetAddress(symbol string) (int, error) {
	if isInt(symbol) {
		return strconv.Atoi(symbol)
	} else if address, isPresent := symbolTable.symbols[symbol]; isPresent {
		return address, nil
	}
	return 0, fmt.Errorf("Can not find symbol in symbol table : : [%s] ", symbol)
}

func initializeSymbolTable() *SymbolTable {
	var symbols = map[string]int{
		"R0":  0,
		"R1":  1,
		"R2":  2,
		"R3":  3,
		"R4":  4,
		"R5":  5,
		"R6":  6,
		"R7":  7,
		"R8":  8,
		"R9":  9,
		"R10": 10,
		"R11": 11,
		"R12": 12,
		"R13": 13,
		"R14": 14,
		"R15": 15,

		"SP":   0,
		"LCL":  1,
		"ARG":  2,
		"THIS": 3,
		"THAT": 4,

		"SCREEN": 16384,
		"KBD":    24576,
	}

	return &SymbolTable{symbols}
}

func (symbolTable *SymbolTable) addLabels(lines []string) {
	var lineIndex int = firstLineIndex
	for _, line := range lines {
		if isLabel(line) {
			symbol := trimLabel(line)
			symbolTable.symbols[symbol] = lineIndex
		} else {
			lineIndex++
		}
	}
}

func (symbolTable *SymbolTable) addVariables(lines []string) {
	var nextFreeRegister int = firstReeRegister
	for _, line := range lines {
		if isAInstruction(line) {
			symbol := trimAInstruction(line)
			if _, isPresent := symbolTable.symbols[symbol]; !isPresent && !isInt(symbol) {
				symbolTable.symbols[symbol] = nextFreeRegister
				nextFreeRegister++
			}
		}
	}
}

func trimLabel(label string) string {
	return strings.Trim(label, "()")
}

func trimAInstruction(aInst string) string {
	return strings.Trim(aInst, "@")
}

func isLabel(text string) bool {
	return strings.HasPrefix(text, "(")
}

func isAInstruction(text string) bool {
	return strings.HasPrefix(text, "@")
}

func isInt(text string) bool {
	if _, err := strconv.Atoi(text); err == nil {
		return true
	}
	return false
}
