package instruction_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zsolt-jakab/nand-to-tetris/assembler/parser/instruction"
)

const (
	invalid string = "invalid"
	opcode  string = "111"
)

func Test_NewC_Should_Raise_Error_When_Destination_Invalid(t *testing.T) {
	expected := "Can not create C instruction with unknown destination: [invalid] "

	actual, err := instruction.NewC(invalid, "M-1", "JGT")

	assert.Equal(t, expected, err.Error())
	assert.Nil(t, actual)
}

func Test_NewC_Should_Raise_Error_When_Computation_Invalid(t *testing.T) {
	expected := "Can not create C instruction with unknown computation: [invalid] "

	actual, err := instruction.NewC("D", invalid, "JGT")

	assert.Equal(t, expected, err.Error())
	assert.Nil(t, actual)
}

func Test_NewC_Should_Raise_Error_When_Jump_Invalid(t *testing.T) {
	expected := "Can not create C instruction with unknown jump: [invalid] "

	actual, err := instruction.NewC("D", "M-1", invalid)

	assert.Equal(t, expected, err.Error())
	assert.Nil(t, actual)
}

func Test_NewC_Destination(t *testing.T) {
	type TestCase struct {
		inputDest      string
		expectedBinary string
	}

	comp := "M-1"
	jump := "JGT"
	compBinary := "1110010"
	jumpBinary := "001"

	tests := []TestCase{
		{inputDest: "", expectedBinary: opcode + compBinary + "000" + jumpBinary},
		{inputDest: "M", expectedBinary: opcode + compBinary + "001" + jumpBinary},
		{inputDest: "D", expectedBinary: opcode + compBinary + "010" + jumpBinary},
		{inputDest: "A", expectedBinary: opcode + compBinary + "100" + jumpBinary},
		{inputDest: "MD", expectedBinary: opcode + compBinary + "011" + jumpBinary},
		{inputDest: "AM", expectedBinary: opcode + compBinary + "101" + jumpBinary},
		{inputDest: "AD", expectedBinary: opcode + compBinary + "110" + jumpBinary},
		{inputDest: "AMD", expectedBinary: opcode + compBinary + "111" + jumpBinary},
	}

	for _, tc := range tests {
		actual, _ := instruction.NewC(tc.inputDest, comp, jump)
		assert.Equal(t, tc.expectedBinary, actual.Binary())
	}
}

func Test_NewC_Computation(t *testing.T) {
	type TestCase struct {
		inputComp      string
		expectedBinary string
	}

	dest := "M"
	jump := "JGT"
	destBinary := "001"
	jumpBinary := "001"

	tests := []TestCase{
		{inputComp: "0", expectedBinary: opcode + "0101010" + destBinary + jumpBinary},
		{inputComp: "1", expectedBinary: opcode + "0111111" + destBinary + jumpBinary},
		{inputComp: "-1", expectedBinary: opcode + "0111010" + destBinary + jumpBinary},
		{inputComp: "D", expectedBinary: opcode + "0001100" + destBinary + jumpBinary},
		{inputComp: "A", expectedBinary: opcode + "0110000" + destBinary + jumpBinary},
		{inputComp: "M", expectedBinary: opcode + "1110000" + destBinary + jumpBinary},
		{inputComp: "!D", expectedBinary: opcode + "0001101" + destBinary + jumpBinary},
		{inputComp: "!A", expectedBinary: opcode + "0110001" + destBinary + jumpBinary},
		{inputComp: "!M", expectedBinary: opcode + "1110001" + destBinary + jumpBinary},
		{inputComp: "-D", expectedBinary: opcode + "0001111" + destBinary + jumpBinary},
		{inputComp: "-A", expectedBinary: opcode + "0110011" + destBinary + jumpBinary},
		{inputComp: "-M", expectedBinary: opcode + "1110011" + destBinary + jumpBinary},
		{inputComp: "D+1", expectedBinary: opcode + "0011111" + destBinary + jumpBinary},
		{inputComp: "A+1", expectedBinary: opcode + "0110111" + destBinary + jumpBinary},
		{inputComp: "M+1", expectedBinary: opcode + "1110111" + destBinary + jumpBinary},
		{inputComp: "D-1", expectedBinary: opcode + "0001110" + destBinary + jumpBinary},
		{inputComp: "A-1", expectedBinary: opcode + "0110010" + destBinary + jumpBinary},
		{inputComp: "M-1", expectedBinary: opcode + "1110010" + destBinary + jumpBinary},
		{inputComp: "D+A", expectedBinary: opcode + "0000010" + destBinary + jumpBinary},
		{inputComp: "D+M", expectedBinary: opcode + "1000010" + destBinary + jumpBinary},
		{inputComp: "D-A", expectedBinary: opcode + "0010011" + destBinary + jumpBinary},
		{inputComp: "D-M", expectedBinary: opcode + "1010011" + destBinary + jumpBinary},
		{inputComp: "A-D", expectedBinary: opcode + "0000111" + destBinary + jumpBinary},
		{inputComp: "M-D", expectedBinary: opcode + "1000111" + destBinary + jumpBinary},
		{inputComp: "D&A", expectedBinary: opcode + "0000000" + destBinary + jumpBinary},
		{inputComp: "D&M", expectedBinary: opcode + "1000000" + destBinary + jumpBinary},
		{inputComp: "D|A", expectedBinary: opcode + "0010101" + destBinary + jumpBinary},
		{inputComp: "D|M", expectedBinary: opcode + "1010101" + destBinary + jumpBinary},
	}

	for _, tc := range tests {
		actual, _ := instruction.NewC(dest, tc.inputComp, jump)
		assert.Equal(t, tc.expectedBinary, actual.Binary())
	}
}

func Test_NewC_Jump(t *testing.T) {
	type TestCase struct {
		inputJump      string
		expectedBinary string
	}

	dest := "M"
	comp := "M-1"
	compBinary := "1110010"
	destBinary := "001"

	tests := []TestCase{
		{inputJump: "JGT", expectedBinary: opcode + compBinary + destBinary + "001"},
		{inputJump: "JEQ", expectedBinary: opcode + compBinary + destBinary + "010"},
		{inputJump: "JGE", expectedBinary: opcode + compBinary + destBinary + "011"},
		{inputJump: "JLT", expectedBinary: opcode + compBinary + destBinary + "100"},
		{inputJump: "JNE", expectedBinary: opcode + compBinary + destBinary + "101"},
		{inputJump: "JLE", expectedBinary: opcode + compBinary + destBinary + "110"},
		{inputJump: "JMP", expectedBinary: opcode + compBinary + destBinary + "111"},
		{inputJump: "", expectedBinary: opcode + compBinary + destBinary + "000"},
	}

	for _, tc := range tests {
		actual, _ := instruction.NewC(dest, comp, tc.inputJump)
		assert.Equal(t, tc.expectedBinary, actual.Binary())
	}
}
