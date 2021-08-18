package compeng

import (
	"github.com/stretchr/testify/assert"
	"github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/analyzer/tokenizer"
	hio "github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/io"
	"io/ioutil"
	"testing"
)

var fileAccessor hio.FileAccessor = &hio.DefaultFileAccessor{FileReader: &hio.DefaultFileReader{}, FileWriter: &hio.DefaultFileWriter{}}


func Test_Compilation_Engine_Square_Main(t *testing.T) {
	var sourceCode = fileAccessor.ReadCode("testdata/square/Main.jack")
	var expected = getFileContentAsString("testdata/square/Main.xml")
	var tokenizer tokenizer.Tokenizer = tokenizer.NewJackTokenizer(sourceCode)
	var compilationEngine CompilationEngine = NewJackCompilationEngine(&tokenizer)

	actual := compilationEngine.CompileClass()

	assert.Equal(t, expected, actual)
}

func Test_Compilation_Engine_Square(t *testing.T) {
	var sourceCode = fileAccessor.ReadCode("testdata/square/Square.jack")
	var expected = getFileContentAsString("testdata/square/Square.xml")
	var tokenizer tokenizer.Tokenizer = tokenizer.NewJackTokenizer(sourceCode)
	var compilationEngine CompilationEngine = NewJackCompilationEngine(&tokenizer)

	actual := compilationEngine.CompileClass()

	assert.Equal(t, expected, actual)
}

func Test_Compilation_Engine_Square_Game(t *testing.T) {
	var sourceCode = fileAccessor.ReadCode("testdata/square/SquareGame.jack")
	var expected = getFileContentAsString("testdata/square/SquareGame.xml")
	var tokenizer tokenizer.Tokenizer = tokenizer.NewJackTokenizer(sourceCode)
	var compilationEngine CompilationEngine = NewJackCompilationEngine(&tokenizer)

	actual := compilationEngine.CompileClass()

	assert.Equal(t, expected, actual)
}

func Test_Compilation_Engine_Array(t *testing.T) {
	var sourceCode = fileAccessor.ReadCode("testdata/array/Main.jack")
	var expected = getFileContentAsString("testdata/array/Main.xml")
	var tokenizer tokenizer.Tokenizer = tokenizer.NewJackTokenizer(sourceCode)
	var compilationEngine CompilationEngine = NewJackCompilationEngine(&tokenizer)

	actual := compilationEngine.CompileClass()

	assert.Equal(t, expected, actual)
}


func getFileContentAsString(fileName string) string {
	return string(getFileContent(fileName))
}

func getFileContent(fileName string) []byte {
	file, _ := ioutil.ReadFile(fileName)
	return file
}
