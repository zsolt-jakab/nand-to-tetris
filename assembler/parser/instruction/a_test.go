package instruction_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zsolt-jakab/nand-to-tetris/assembler/parser/instruction"
)

const (
	lastValidAddress = 32767
)

func Test_NewC_Should_Raise_Error_When_Address_Negative(t *testing.T) {
	expected := "Can not create A instruction with negative address value: [-1] "

	actual, err := instruction.NewA(-1)

	assert.Equal(t, expected, err.Error())
	assert.Nil(t, actual)
}

func Test_NewC_Should_Raise_Error_When_Address_Bigger_Than_Limit(t *testing.T) {
	expected := "Can not create A instruction with bigger than 32767 address value: [32768] "

	actual, err := instruction.NewA(lastValidAddress + 1)

	assert.Equal(t, expected, err.Error())
	assert.Nil(t, actual)
}

func Test_NewA(t *testing.T) {
	type TestCase struct {
		inputDecimalAddress int
		expectedBinary      string
	}

	testCases := []TestCase{
		{inputDecimalAddress: 0, expectedBinary: "0000000000000000"},
		{inputDecimalAddress: 21, expectedBinary: "0000000000010101"},
		{inputDecimalAddress: lastValidAddress, expectedBinary: "0111111111111111"},
	}

	for _, testCase := range testCases {
		actual, _ := instruction.NewA(testCase.inputDecimalAddress)
		assert.Equal(t, testCase.expectedBinary, actual.Binary())
	}
}
