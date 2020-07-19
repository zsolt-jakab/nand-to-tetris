// +build integration

package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testFileName             = "testdata/test"
	testHackFileName         = testFileName + ".hack"
	verificationHackFileName = "testdata/verification.hack"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Remove(testHackFileName)
	
	os.Exit(exitVal)
}

func Test_Do_Main_Integration(t *testing.T) {
	doMain(testFileName)

	assert.Equal(t, getFileContent(verificationHackFileName), getFileContent(testHackFileName))
}

func Test_ReadCodeLines_Panic_When_Wrong_FileName_Integration(t *testing.T) {
	action := func() { doMain("/wrong/directory/test") }

	assert.PanicsWithError(t, "open /wrong/directory/test.asm: no such file or directory", action)
}

func getFileContent(fileName string) string {
	file, _ := ioutil.ReadFile(fileName)
	return string(file)
}
