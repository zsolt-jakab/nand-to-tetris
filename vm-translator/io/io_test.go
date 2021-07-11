package io_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	vmio "github.com/zsolt-jakab/nand-to-tetris/vm-translator/io"
)

const (
	baseTestFileName = "test"
	testFileName     = "testdata/test"
	testAsmFileName  = testFileName + ".asm"
	testVmFileName   = testFileName + ".vm"
	newLine          = "\n"
)

type ReaderMock struct {
	mock.Mock
}

type WriterMock struct {
	mock.Mock
}

type AccessMock struct {
	mock.Mock
}

type FileInfoMock struct {
	IsDirectory bool
}

func (fim FileInfoMock) Name() string       { return "" }
func (fim FileInfoMock) Size() int64        { return int64(8) }
func (fim FileInfoMock) Mode() os.FileMode  { return os.ModePerm }
func (fim FileInfoMock) ModTime() time.Time { return time.Now() }
func (fim FileInfoMock) IsDir() bool        { return fim.IsDirectory }
func (fim FileInfoMock) Sys() interface{}   { return nil }

func (rm *ReaderMock) Read(fileName string) ([]byte, error) {
	args := rm.Called(fileName)

	return args.Get(0).([]byte), args.Error(1)
}

func (wm *WriterMock) Write(name string, data []byte) error {
	args := wm.Called(name, data)

	return args.Error(0)
}

func (am *AccessMock) Stat(fileName string) (os.FileInfo, error) {
	args := am.Called(fileName)

	return args.Get(0).(os.FileInfo), args.Error(1)
}

func Test_ReadCodeLines(t *testing.T) {
	expectedCodeLines := getCodeLines("testdata/vm/simple/expected.txt")
	readerMockResponse := getFileContent("testdata/vm/simple/mock_response.txt")
	readerMock := stubReader(readerMockResponse)
	accessMock := stubAccess()
	sut := vmio.VMTranslatorFileAccessor{FileReader: readerMock, FileWriter: &vmio.DefaultFileWriter{}, FileAccess: accessMock}

	actualCodeLines := sut.ReadSourceLines(testVmFileName)

	assert.Equal(t, expectedCodeLines, actualCodeLines)
	readerMock.AssertExpectations(t)
}

func Test_ReadCodeLines_When_Instruction_Comments(t *testing.T) {
	expectedCodeLineIndexes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}
	expectedCodeLines := getCodeLinesWithIndexes("testdata/vm/instruction_comments/expected.txt", expectedCodeLineIndexes)
	readerMockResponse := getFileContent("testdata/vm/instruction_comments/mock_response.txt")
	readerMock := new(ReaderMock)
	readerMock.On("Read", testVmFileName).Return(readerMockResponse, nil)
	accessMock := stubAccess()

	sut := vmio.VMTranslatorFileAccessor{FileReader: readerMock, FileWriter: &vmio.DefaultFileWriter{}, FileAccess: accessMock}

	actualCodeLines := sut.ReadSourceLines(testVmFileName)

	assert.Equal(t, expectedCodeLines, actualCodeLines)
	readerMock.AssertExpectations(t)
}

func Test_ReadCodeLines_When_Empty_Lines(t *testing.T) {
	expectedCodeLineIndexes := []int{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	expectedCodeLines := getCodeLinesWithIndexes("testdata/vm/empty_lines/expected.txt", expectedCodeLineIndexes)
	readerMockResponse := getFileContent("testdata/vm/empty_lines/mock_response.txt")
	readerMock := stubReader(readerMockResponse)
	accessMock := stubAccess()

	sut := vmio.VMTranslatorFileAccessor{FileReader: readerMock, FileWriter: &vmio.DefaultFileWriter{}, FileAccess: accessMock}

	actualCodeLines := sut.ReadSourceLines(testVmFileName)

	assert.Equal(t, expectedCodeLines, actualCodeLines)
	readerMock.AssertExpectations(t)
}

func Test_ReadCodeLines_When_Comment_Lines(t *testing.T) {
	expectedCodeLineIndexes := []int{9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 24, 25, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
	expectedCodeLines := getCodeLinesWithIndexes("testdata/vm/comment_lines/expected.txt", expectedCodeLineIndexes)
	readerMockResponse := getFileContent("testdata/vm/comment_lines/mock_response.txt")
	readerMock := stubReader(readerMockResponse)
	accessMock := stubAccess()
	sut := vmio.VMTranslatorFileAccessor{FileReader: readerMock, FileWriter: &vmio.DefaultFileWriter{}, FileAccess: accessMock}

	actualCodeLines := sut.ReadSourceLines(testVmFileName)

	assert.Equal(t, expectedCodeLines, actualCodeLines)
	readerMock.AssertExpectations(t)
}

func Test_ReadCodeLines_Panic_When_Read_Error(t *testing.T) {
	readerMock := new(ReaderMock)
	readerMock.On("Read", testVmFileName).Return([]byte{}, fmt.Errorf("Error message"))
	accessMock := stubAccess()

	sut := vmio.VMTranslatorFileAccessor{FileReader: readerMock, FileWriter: &vmio.DefaultFileWriter{}, FileAccess: accessMock}

	action := func() { sut.ReadSourceLines(testVmFileName) }

	assert.PanicsWithError(t, "Error message", action)
}

func Test_CreateHackFile(t *testing.T) {
	writerMock := new(WriterMock)
	writerMock.On("Write", testAsmFileName, []byte("line 1\nline 2\nline 3\n")).Return(nil)
	linesToSave := []string{"line 1", "line 2", "line 3"}
	accessMock := stubAccess()

	sut := vmio.VMTranslatorFileAccessor{FileReader: &vmio.DefaultFileReader{}, FileWriter: writerMock, FileAccess: accessMock}

	sut.CreateTargetFile(testVmFileName, linesToSave)

	writerMock.AssertExpectations(t)
}

func Test_CreateHackFile_Panic_When_Write_Error(t *testing.T) {
	writerMock := new(WriterMock)
	writerMock.On("Write", testAsmFileName, []byte("line 1\nline 2\nline 3\n")).Return(fmt.Errorf("Error message"))
	linesToSave := []string{"line 1", "line 2", "line 3"}
	accessMock := stubAccess()

	sut := vmio.VMTranslatorFileAccessor{FileReader: &vmio.DefaultFileReader{}, FileWriter: writerMock, FileAccess: accessMock}

	action := func() { sut.CreateTargetFile(testVmFileName, linesToSave) }

	assert.PanicsWithError(t, "Error message", action)
	writerMock.AssertExpectations(t)
}

func stubReader(readerMockResponse []byte) *ReaderMock {
	readerMock := new(ReaderMock)
	readerMock.On("Read", testVmFileName).Return(readerMockResponse, nil)
	return readerMock
}

func stubAccess() *AccessMock {
	fileInfoMock := &FileInfoMock{false}

	accessMock := new(AccessMock)
	accessMock.On("Stat", testVmFileName).Return(fileInfoMock, nil)
	return accessMock
}

func getCodeLines(path string) []vmio.CodeLine {
	rawLines := strings.Fields(getFileContentAsString(path))
	var codeLines []vmio.CodeLine
	for i, rawLine := range rawLines {
		codeLines = append(codeLines, vmio.CodeLine{FileName: baseTestFileName, Content: rawLine, Index: i + 1})
	}
	return codeLines
}

func getCodeLinesWithIndexes(path string, indexes []int) []vmio.CodeLine {
	rawLines := strings.Fields(getFileContentAsString(path))
	var codeLines []vmio.CodeLine
	for i, rawLine := range rawLines {
		codeLines = append(codeLines, vmio.CodeLine{FileName: baseTestFileName, Content: rawLine, Index: indexes[i]})
	}
	return codeLines
}

func getFileContentAsString(fileName string) string {
	return string(getFileContent(fileName))
}

func getFileContent(fileName string) []byte {
	file, _ := ioutil.ReadFile(fileName)
	return file
}
