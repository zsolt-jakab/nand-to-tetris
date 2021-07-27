// +build integration

package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

const (
	testFileName             = "testdata"
	testHackFileName         = testFileName + ".hack"
)

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Remove("testdata/array/Main.xml")
	os.Remove("testdata/square/Main.xml")
	os.Remove("testdata/square/Square.xml")
	os.Remove("testdata/square/SquareGame.xml")

	os.Exit(exitVal)
}

func Test_Do_Main_Integration(t *testing.T) {
	doMain(testFileName)

	assert.Equal(t, getFileContent("testdata/array/MainVerification.xml"), getFileContent("testdata/array/Main.xml"))
	assert.Equal(t, getFileContent("testdata/square/MainVerification.xml"), getFileContent("testdata/square/Main.xml"))
	assert.Equal(t, getFileContent("testdata/square/SquareVerification.xml"), getFileContent("testdata/square/Square.xml"))
	assert.Equal(t, getFileContent("testdata/square/SquareGameVerification.xml"), getFileContent("testdata/square/SquareGame.xml"))
}

func getFileContent(fileName string) string {
	file, _ := ioutil.ReadFile(fileName)
	return string(file)
}