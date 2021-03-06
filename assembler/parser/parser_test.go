  package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zsolt-jakab/nand-to-tetris/assembler/parser"
)

func Test_Translate_A_Instruction(t *testing.T) {
	expected := []string{"0000000000000000", "0000000000010000", "0000000000010001", "0000000000010000", "0000000000010001"}
	lines := []string{"@0", "@firstVar", "@secondVar", "@firstVar", "@secondVar"}
	lineIndexes := []int{1, 2, 3, 4, 5}

	actual := parser.Translate(lines, lineIndexes)

	assert.Equal(t, expected, actual)
}

func Test_Translate_A_Instruction_With_Label(t *testing.T) {
	expected := []string{"0000000000000000", "0000000000000010", "0000000000010000"}
	lines := []string{"@0", "@LABEL", "(LABEL)", "@variable"}
	lineIndexes := []int{1, 2, 3, 4}

	actual := parser.Translate(lines, lineIndexes)

	assert.Equal(t, expected, actual)
}

func Test_Translate_Invalid_A_Instruction_Panics(t *testing.T) {
	panicMessage := "Error: [Can not create A instruction with bigger than 32767 address value: [32768] ] in code: [@32768] in line: [1]"
	lines := []string{"@32768"}
	lineIndexes := []int{1}

	action := func() { parser.Translate(lines, lineIndexes) }

	assert.PanicsWithValue(t, panicMessage, action)
}

func Test_Translate_C_Instruction(t *testing.T) {
	expected := []string{"1111000010001111", "1111110111001000", "1110001100000001", "1111000010010000"}
	lines := []string{"M=D+M;JMP", "M=M+1", "D;JGT", "D=D+M"}
	lineIndexes := []int{1, 2, 3, 4}

	actual := parser.Translate(lines, lineIndexes)

	assert.Equal(t, expected, actual)
}

func Test_Translate_Invalid_C_Panics(t *testing.T) {
	panicMessage := "Error: [Can not create C instruction with unknown computation: [inv-comp] ] in code: [M=inv-comp;JGT] in line: [1]"
	lines := []string{"M=inv-comp;JGT"}
	lineIndexes := []int{1}

	action := func() { parser.Translate(lines, lineIndexes) }

	assert.PanicsWithValue(t, panicMessage, action)
}

func Test_Translate_Unknown_Instruction_Panics(t *testing.T) {
	panicMessage := "Error: [Can not create C instruction with unknown computation: [Custom Instruction] ] in code: [Custom Instruction] in line: [2]"
	lines := []string{"@0", "Custom Instruction"}
	lineIndexes := []int{1, 2}

	action := func() { parser.Translate(lines, lineIndexes) }

	assert.PanicsWithValue(t, panicMessage, action)
}
