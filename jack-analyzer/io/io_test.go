package io_test

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	hio "github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/io"
)

const (
	testFileName         = "testdata/test"
	testHackFileName     = testFileName + ".hack"
	testAssemblyFileName = testFileName + ".asm"
)

type ReaderMock struct {
	mock.Mock
}

type WriterMock struct {
	mock.Mock
}

func (rm *ReaderMock) Read(fileName string) ([]byte, error) {
	args := rm.Called(fileName)

	return args.Get(0).([]byte), args.Error(1)
}

func (wm *WriterMock) Write(name string, data []byte) error {
	args := wm.Called(name, data)

	return args.Error(0)
}

func Test_ReadCodeLines(t *testing.T) {
	expectedCodeLineIndexes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}
	expectedCodeLines := getFileContentAsLines("testdata/asm/simple/expected.txt")
	readerMockResponse := getFileContent("testdata/asm/simple/mock_response.txt")
	readerMock := stubReader(readerMockResponse)
	sut := hio.DefaultFileAccessor{readerMock, &hio.DefaultFileWriter{}}
	actualCodeLines, actualCodeLineIndexes := sut.ReadCodeAsLines(testAssemblyFileName)

	assert.Equal(t, expectedCodeLines, actualCodeLines)
	assert.Equal(t, expectedCodeLineIndexes, actualCodeLineIndexes)
	readerMock.AssertExpectations(t)
}

func Test_ReadCodeLines_When_Instruction_Comments(t *testing.T) {
	expectedCodeLineIndexes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}
	expectedCodeLines := getFileContentAsLines("testdata/asm/instruction_comments/expected.txt")
	readerMockResponse := getFileContent("testdata/asm/instruction_comments/mock_response.txt")
	readerMock := new(ReaderMock)
	readerMock.On("Read", testAssemblyFileName).Return(readerMockResponse, nil)
	sut := hio.DefaultFileAccessor{readerMock, &hio.DefaultFileWriter{}}

	actualCodeLines, actualCodeLineIndexes := sut.ReadCodeAsLines(testAssemblyFileName)

	assert.Equal(t, expectedCodeLines, actualCodeLines)
	assert.Equal(t, expectedCodeLineIndexes, actualCodeLineIndexes)
	readerMock.AssertExpectations(t)
}

func Test_ReadCodeLines_When_Empty_Lines(t *testing.T) {
	expectedCodeLineIndexes := []int{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	expectedCodeLines := getFileContentAsLines("testdata/asm/empty_lines/expected.txt")
	readerMockResponse := getFileContent("testdata/asm/empty_lines/mock_response.txt")
	readerMock := stubReader(readerMockResponse)
	sut := hio.DefaultFileAccessor{readerMock, &hio.DefaultFileWriter{}}

	actualCodeLines, actualCodeLineIndexes := sut.ReadCodeAsLines(testAssemblyFileName)

	assert.Equal(t, expectedCodeLines, actualCodeLines)
	assert.Equal(t, expectedCodeLineIndexes, actualCodeLineIndexes)
	readerMock.AssertExpectations(t)
}

func Test_ReadCodeLines_When_Comment_Lines(t *testing.T) {
	expectedCodeLineIndexes := []int{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	expectedCodeLines := getFileContentAsLines("testdata/asm/comment_lines/expected.txt")
	readerMockResponse := getFileContent("testdata/asm/comment_lines/mock_response.txt")
	readerMock := stubReader(readerMockResponse)
	sut := hio.DefaultFileAccessor{readerMock, &hio.DefaultFileWriter{}}

	actualCodeLines, actualCodeLineIndexes := sut.ReadCodeAsLines(testAssemblyFileName)

	assert.Equal(t, expectedCodeLines, actualCodeLines)
	assert.Equal(t, expectedCodeLineIndexes, actualCodeLineIndexes)
	readerMock.AssertExpectations(t)
}

func Test_ReadCodeLines_Panic_When_Read_Error(t *testing.T) {
	readerMock := new(ReaderMock)
	readerMock.On("Read", testAssemblyFileName).Return([]byte{}, fmt.Errorf("Error message "))

	sut := hio.DefaultFileAccessor{readerMock, &hio.DefaultFileWriter{}}

	action := func() { sut.ReadCodeAsLines(testAssemblyFileName) }

	assert.PanicsWithError(t, "Error message ", action)
}

func Test_CreateFile(t *testing.T) {
	writerMock := new(WriterMock)
	writerMock.On("Write", testHackFileName, []byte("line 1\nline 2\nline 3\n")).Return(nil)
	linesToSave := []string{"line 1", "line 2", "line 3"}
	sut := hio.DefaultFileAccessor{&hio.DefaultFileReader{}, writerMock}

	sut.CreateFile(testHackFileName, linesToSave)

	writerMock.AssertExpectations(t)
}

func Test_CreateFile_Panic_When_Write_Error(t *testing.T) {
	writerMock := new(WriterMock)
	writerMock.On("Write", testHackFileName, []byte("line 1\nline 2\nline 3\n")).Return(fmt.Errorf("Error message"))
	linesToSave := []string{"line 1", "line 2", "line 3"}
	sut := hio.DefaultFileAccessor{&hio.DefaultFileReader{}, writerMock}

	action := func() { sut.CreateFile(testHackFileName, linesToSave) }

	assert.PanicsWithError(t, "Error message", action)
	writerMock.AssertExpectations(t)
}

func stubReader(readerMockResponse []byte) *ReaderMock {
	readerMock := new(ReaderMock)
	readerMock.On("Read", testAssemblyFileName).Return(readerMockResponse, nil)
	return readerMock
}

func getFileContentAsLines(fileName string) []string {
	return strings.Fields(getFileContentAsString(fileName))
}

func getFileContentAsString(fileName string) string {
	return string(getFileContent(fileName))
}

func getFileContent(fileName string) []byte {
	file, _ := ioutil.ReadFile(fileName)
	return file
}
