package tokenizer

import (
	"github.com/stretchr/testify/assert"
	"github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/analyzer"
	hio "github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/io"
	"html"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

var fileAccessor hio.FileAccessor = &hio.DefaultFileAccessor{FileReader: &hio.DefaultFileReader{}, FileWriter: &hio.DefaultFileWriter{}}

func Test_Tokenizer_Array(t *testing.T) {
	var sourceCode = fileAccessor.ReadCode("testdata/array/Main.jack")
	var expected = getFileContentAsString("testdata/array/MainT.xml")
	var tokenizer analyzer.Tokenizer = NewJackTokenizer(sourceCode)

	actual := tokenize(tokenizer)

	assert.Equal(t, expected, actual)
}

func Test_Tokenizer_Square_Main(t *testing.T) {
	var sourceCode = fileAccessor.ReadCode("testdata/square/Main.jack")
	var expected = getFileContentAsString("testdata/square/MainT.xml")
	var tokenizer analyzer.Tokenizer = NewJackTokenizer(sourceCode)

	actual := tokenize(tokenizer)

	assert.Equal(t, expected, actual)
}

func Test_Tokenizer_Square(t *testing.T) {
	var sourceCode = fileAccessor.ReadCode("testdata/square/Square.jack")
	var expected = getFileContentAsString("testdata/square/SquareT.xml")
	var tokenizer analyzer.Tokenizer = NewJackTokenizer(sourceCode)

	actual := tokenize(tokenizer)

	assert.Equal(t, expected, actual)
}

func Test_Tokenizer_Square_Game(t *testing.T) {
	var sourceCode = fileAccessor.ReadCode("testdata/square/SquareGame.jack")
	var expected = getFileContentAsString("testdata/square/SquareGameT.xml")
	var tokenizer analyzer.Tokenizer = NewJackTokenizer(sourceCode)

	actual := tokenize(tokenizer)

	assert.Equal(t, expected, actual)
}


func tokenize(tokenizer analyzer.Tokenizer) string {
	var sb strings.Builder
	sb.WriteString("<tokens>")
	sb.WriteString("\r\n")
	for tokenizer.HasMoreTokens() {
		tokenizer.Advance()

		sb.WriteString("<" + tokenizer.TokenType().String() + ">")
		sb.WriteString(" " + getTokenValue(tokenizer) + " ")
		sb.WriteString("</" + tokenizer.TokenType().String() + ">")
		sb.WriteString("\r\n")
	}
	sb.WriteString("</tokens>")
	sb.WriteString("\r\n")
	var actual = sb.String()
	return actual
}

func getTokenValue(tokenizer analyzer.Tokenizer) string {
	switch tokenizer.TokenType() {
	case analyzer.Symbol:
		return html.EscapeString(string(tokenizer.Symbol()))
	case analyzer.Keyword:
		return tokenizer.KeyWord()
	case analyzer.Identifier:
		return tokenizer.Identifier()
	case analyzer.IntConst:
		return strconv.Itoa(tokenizer.IntVal())
	case analyzer.StringConst:
		return tokenizer.StringVal()
	default:
		return "Unknown"
	}
}

func getFileContentAsString(fileName string) string {
	return string(getFileContent(fileName))
}

func getFileContent(fileName string) []byte {
	file, _ := ioutil.ReadFile(fileName)
	return file
}
