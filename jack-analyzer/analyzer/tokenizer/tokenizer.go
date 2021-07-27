package tokenizer

import (
	"strconv"
	"unicode"
)

const Quote = '"'

type JackTokenizer struct {
	sourceCode       string
	currentCharIndex int

	tokenType  TokenType
	keyWord    string
	symbol     uint8
	identifier string
	intVal     int
	stringVal  string
}

func NewJackTokenizer(sourceCode string) *JackTokenizer {
	return &JackTokenizer{
		sourceCode:       sourceCode,
		currentCharIndex: 0,
		tokenType:        UnknownToken,
	}
}

func (jT *JackTokenizer) KeyWord() string {
	return jT.keyWord
}

func (jT *JackTokenizer) Symbol() uint8 {
	return jT.symbol
}

func (jT *JackTokenizer) Identifier() string {
	return jT.identifier
}

func (jT *JackTokenizer) IntVal() int {
	return jT.intVal
}

func (jT *JackTokenizer) StringVal() string {
	return jT.stringVal
}

func (jT *JackTokenizer) HasMoreTokens() bool {
	return jT.currentCharIndex < jT.lastSourceCodeIndex()
}

func (jT *JackTokenizer) Advance() {
	jT.skipSpaces()
	if jT.isCurrentSymbol() {
		jT.advanceSymbol()
	} else if jT.isCurrentStringConstant() {
		jT.advanceStringConstant()
	} else if jT.isCurrentIntConstant() {
		jT.advanceIntConstant()
	} else {
		jT.advanceIdentifierOrKeyword()
	}
}

func (jT *JackTokenizer) advanceIdentifierOrKeyword() {
	startOfNexToken := jT.getEndOfWord() + 1
	currentWord := jT.sourceCode[jT.currentCharIndex:startOfNexToken]
	if Keywords[currentWord] {
		jT.tokenType = Keyword
		jT.keyWord = currentWord
	} else {
		jT.tokenType = Identifier
		jT.identifier = currentWord
	}
	jT.currentCharIndex = startOfNexToken
}

func (jT *JackTokenizer) TokenType() TokenType {
	return jT.tokenType
}

func (jT *JackTokenizer) lastSourceCodeIndex() int {
	return len(jT.sourceCode) - 1
}

func (jT *JackTokenizer) isCurrentStringConstant() bool {
	return jT.getCurrentChar() == Quote
}

func (jT *JackTokenizer) isCurrentIntConstant() bool {
	return unicode.IsDigit(rune(jT.sourceCode[jT.currentCharIndex]))
}

func (jT *JackTokenizer) advanceSymbol() {
	jT.tokenType = Symbol
	jT.symbol = jT.getCurrentChar()
	jT.currentCharIndex++
}

func (jT *JackTokenizer) isCurrentSymbol() bool {
	var currentChar = jT.getCurrentChar()

	return Symbols[currentChar]
}

func (jT *JackTokenizer) getCurrentChar() uint8 {
	return jT.sourceCode[jT.currentCharIndex]
}

func (jT *JackTokenizer) advanceStringConstant() {
	jT.tokenType = StringConst
	endQuoteIndex := jT.getEndQuoteIndex()
	jT.stringVal = jT.sourceCode[jT.currentCharIndex+1 : endQuoteIndex]

	jT.currentCharIndex = endQuoteIndex + 1
}

func (jT *JackTokenizer) advanceIntConstant() {
	jT.tokenType = IntConst
	startOfNextToken := jT.getEndOfInt() + 1

	jT.intVal, _ = strconv.Atoi(jT.sourceCode[jT.currentCharIndex:startOfNextToken])
	jT.currentCharIndex = startOfNextToken
}

func (jT *JackTokenizer) getEndQuoteIndex() int {
	endQuoteIndex := jT.currentCharIndex + 1
	for jT.sourceCode[endQuoteIndex] != Quote {
		endQuoteIndex++
	}
	return endQuoteIndex
}

func (jT *JackTokenizer) getEndOfInt() int {
	startOfNextToken := jT.currentCharIndex + 1
	for unicode.IsDigit(rune(jT.sourceCode[startOfNextToken])) {
		startOfNextToken++
	}
	return startOfNextToken - 1
}

func (jT *JackTokenizer) getEndOfWord() int {
	endOfCurrentWord := jT.currentCharIndex + 1
	for unicode.IsDigit(rune(jT.sourceCode[endOfCurrentWord])) ||
		unicode.IsLetter(rune(jT.sourceCode[endOfCurrentWord])) ||
		jT.sourceCode[endOfCurrentWord] == '_' {
		endOfCurrentWord++
	}
	return endOfCurrentWord - 1
}

func (jT *JackTokenizer) skipSpaces() {
	nextNotSpace := jT.currentCharIndex
	for unicode.IsSpace(rune(jT.sourceCode[nextNotSpace])) {
		nextNotSpace++
	}
	jT.currentCharIndex = nextNotSpace
}
