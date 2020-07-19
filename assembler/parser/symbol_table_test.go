package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zsolt-jakab/nand-to-tetris/assembler/parser"
)

func Test_NewSymbolTable_Pre_DefinedSymbols(t *testing.T) {
	type TestCase struct {
		symbol  string
		address int
	}

	testCases := []TestCase{
		{symbol: "R0", address: 0},
		{symbol: "R1", address: 1},
		{symbol: "R2", address: 2},
		{symbol: "R3", address: 3},
		{symbol: "R4", address: 4},
		{symbol: "R5", address: 5},
		{symbol: "R6", address: 6},
		{symbol: "R7", address: 7},
		{symbol: "R8", address: 8},
		{symbol: "R9", address: 9},
		{symbol: "R10", address: 10},
		{symbol: "R11", address: 11},
		{symbol: "R12", address: 12},
		{symbol: "R13", address: 13},
		{symbol: "R14", address: 14},
		{symbol: "R15", address: 15},
		{symbol: "SP", address: 0},
		{symbol: "LCL", address: 1},
		{symbol: "ARG", address: 2},
		{symbol: "THIS", address: 3},
		{symbol: "THAT", address: 4},
		{symbol: "SCREEN", address: 16384},
		{symbol: "KBD", address: 24576},
	}

	for _, testCase := range testCases {
		symbolTable := parser.NewSymbolTable([]string{})
		actual, _ := symbolTable.GetAddress(testCase.symbol)
		assert.Equal(t, testCase.address, actual)
	}
}

func Test_Label_When_First_Instruction(t *testing.T) {
	lines := []string{"(label)", "not a label"}

	symbolTable := parser.NewSymbolTable(lines)

	actual, _ := symbolTable.GetAddress("label")
	assert.Equal(t, 0, actual)
}

func Test_Label_When_There_Are_Other_Type_Of_Instructions(t *testing.T) {
	lines := []string{"not a label", "@variable", "(label)", "not a label"}

	symbolTable := parser.NewSymbolTable(lines)

	actual, _ := symbolTable.GetAddress("label")
	assert.Equal(t, 2, actual)
}

func Test_Variable(t *testing.T) {
	lines := []string{"@variable"}

	symbolTable := parser.NewSymbolTable(lines)

	actual, _ := symbolTable.GetAddress("variable")
	assert.Equal(t, 16, actual)
}

func Test_Variable_When_Is_int(t *testing.T) {
	lines := []string{"@1"}

	symbolTable := parser.NewSymbolTable(lines)

	actual, _ := symbolTable.GetAddress("1")
	assert.Equal(t, 1, actual)
}

func Test_Variable_When_Is_Repeating_Should_Not_Update_Value(t *testing.T) {
	lines := []string{"@variable", "@variable"}

	symbolTable := parser.NewSymbolTable(lines)

	actual, _ := symbolTable.GetAddress("variable")
	assert.Equal(t, 16, actual)
}

func Test_Variable_When_There_Are_Other_Type_Of_Instructions(t *testing.T) {
	lines := []string{"not a variable", "(label)", "@variable", "not a variable"}

	symbolTable := parser.NewSymbolTable(lines)

	actual, _ := symbolTable.GetAddress("variable")
	assert.Equal(t, 16, actual)
}

func Test_Multiple_Variables_Incrementing_Value(t *testing.T) {
	lines := []string{"@variable1", "@variable2", "@variable3"}

	symbolTable := parser.NewSymbolTable(lines)

	variable1, _ := symbolTable.GetAddress("variable1")
	variable2, _ := symbolTable.GetAddress("variable2")
	variable3, _ := symbolTable.GetAddress("variable3")
	assert.Equal(t, 16, variable1)
	assert.Equal(t, 17, variable2)
	assert.Equal(t, 18, variable3)
}
