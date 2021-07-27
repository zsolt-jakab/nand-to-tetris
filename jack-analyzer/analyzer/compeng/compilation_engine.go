package compeng

import (
	"github.com/zsolt-jakab/nand-to-tetris/jack-analyzer/analyzer"
	"html"
	"strings"
)

var Operands = map[uint8]bool{
	'+': true,
	'-': true,
	'*': true,
	'/': true,
	'&': true,
	'|': true,
	'<': true,
	'>': true,
	'=': true,
}

type JackCompilationEngine struct {
	tokenizer        analyzer.Tokenizer
	doubleSpaceCount int
	sb               strings.Builder
}

func NewJackCompilationEngine(tokenizer *analyzer.Tokenizer) *JackCompilationEngine {
	return &JackCompilationEngine{
		tokenizer:        *tokenizer,
		doubleSpaceCount: 0,
	}
}

func (jCE *JackCompilationEngine) CompileClass() string {
	jCE.writeLine("<class>")
	jCE.doubleSpaceCount++
	jCE.tokenizer.Advance()

	jCE.writeLine("<keyword> class </keyword>")
	jCE.tokenizer.Advance()
	jCE.writeLine("<identifier> " + jCE.tokenizer.Identifier() + " </identifier>")
	jCE.tokenizer.Advance()
	jCE.writeLine("<symbol> { </symbol>")
	jCE.tokenizer.Advance()
	for jCE.isClassVar() {
		jCE.CompileClassVarDec()
	}
	for jCE.isSubroutineDec() {
		jCE.CompileSubroutineDec()
	}

	jCE.writeLine("<symbol> } </symbol>")
	jCE.doubleSpaceCount--
	jCE.writeLine("</class>")
	return jCE.sb.String()
}

func (jCE *JackCompilationEngine) CompileSubroutineDec() {
	jCE.writeLine("<subroutineDec>")
	jCE.doubleSpaceCount++
	jCE.compileKeywordOrIdentifiers()

	jCE.writeLine("<symbol> ( </symbol>")
	jCE.tokenizer.Advance()

	jCE.CompileParameterList()

	jCE.writeLine("<symbol> ) </symbol>")

	jCE.tokenizer.Advance()
	jCE.CompileSubroutineBody()

	jCE.doubleSpaceCount--
	jCE.writeLine("</subroutineDec>")
}

func (jCE *JackCompilationEngine) CompileSubroutineBody() {
	jCE.writeLine("<subroutineBody>")
	jCE.doubleSpaceCount++
	jCE.writeLine("<symbol> { </symbol>")
	jCE.tokenizer.Advance()

	for jCE.tokenizer.TokenType() == analyzer.Keyword && jCE.tokenizer.KeyWord() == "var" {
		jCE.CompileVarDec()
	}

	jCE.CompileStatements()

	jCE.writeLine("<symbol> } </symbol>")
	jCE.tokenizer.Advance()

	jCE.doubleSpaceCount--
	jCE.writeLine("</subroutineBody>")
}

func (jCE *JackCompilationEngine) CompileStatements() {
	jCE.writeLine("<statements>")
	jCE.doubleSpaceCount++
	for jCE.isStatement() {
		switch jCE.tokenizer.KeyWord() {
		case "let":
			jCE.CompileLet()
		case "if":
			jCE.CompileIf()
		case "while":
			jCE.CompileWhile()
		case "do":
			jCE.CompileDo()
		case "return":
			jCE.CompileReturn()
		}
	}
	jCE.doubleSpaceCount--
	jCE.writeLine("</statements>")
}

func (jCE *JackCompilationEngine) CompileDo() {
	jCE.writeLine("<doStatement>")
	jCE.doubleSpaceCount++

	jCE.writeLine("<keyword> " + jCE.tokenizer.KeyWord() + " </keyword>")
	jCE.tokenizer.Advance()

	jCE.CompileSubroutineCall()

	jCE.writeLine("<symbol> ; </symbol>")
	jCE.tokenizer.Advance()

	jCE.doubleSpaceCount--
	jCE.writeLine("</doStatement>")
}

func (jCE *JackCompilationEngine) CompileSubroutineCall() {
	jCE.writeLine("<identifier> " + jCE.tokenizer.Identifier() + " </identifier>")
	jCE.tokenizer.Advance()
	if jCE.tokenizer.TokenType() == analyzer.Symbol && jCE.tokenizer.Symbol() == '.' {
		jCE.writeLine("<symbol> . </symbol>")
		jCE.tokenizer.Advance()
		jCE.writeLine("<identifier> " + jCE.tokenizer.Identifier() + " </identifier>")
		jCE.tokenizer.Advance()
	}

	jCE.compileExpressionList()

}

func (jCE *JackCompilationEngine) compileExpressionList() {
	jCE.writeLine("<symbol> ( </symbol>")
	jCE.writeLine("<expressionList>")
	jCE.doubleSpaceCount++

	jCE.tokenizer.Advance()
	if jCE.tokenizer.TokenType() != analyzer.Symbol || jCE.tokenizer.Symbol() != ')' {
		jCE.CompileExpression()
	}
	for jCE.tokenizer.TokenType() == analyzer.Symbol && jCE.tokenizer.Symbol() == ',' {
		jCE.writeLine("<symbol> , </symbol>")
		jCE.tokenizer.Advance()

		jCE.CompileExpression()
	}
	jCE.doubleSpaceCount--
	jCE.writeLine("</expressionList>")
	jCE.writeLine("<symbol> ) </symbol>")
	jCE.tokenizer.Advance()
}

func (jCE *JackCompilationEngine) CompileReturn() {
	jCE.writeLine("<returnStatement>")
	jCE.doubleSpaceCount++

	jCE.writeLine("<keyword> " + jCE.tokenizer.KeyWord() + " </keyword>")
	jCE.tokenizer.Advance()

	if jCE.tokenizer.TokenType() != analyzer.Symbol || jCE.tokenizer.Symbol() != ';' {
		jCE.CompileExpression()
	}
	jCE.writeLine("<symbol> ; </symbol>")
	jCE.tokenizer.Advance()

	jCE.doubleSpaceCount--
	jCE.writeLine("</returnStatement>")
}

func (jCE *JackCompilationEngine) CompileWhile() {
	jCE.writeLine("<whileStatement>")
	jCE.doubleSpaceCount++

	jCE.writeLine("<keyword> " + jCE.tokenizer.KeyWord() + " </keyword>")
	jCE.tokenizer.Advance()

	jCE.writeLine("<symbol> ( </symbol>")
	jCE.tokenizer.Advance()

	jCE.CompileExpression()

	jCE.writeLine("<symbol> ) </symbol>")
	jCE.tokenizer.Advance()

	jCE.writeLine("<symbol> { </symbol>")
	jCE.tokenizer.Advance()

	jCE.CompileStatements()

	jCE.writeLine("<symbol> } </symbol>")

	jCE.tokenizer.Advance()

	jCE.doubleSpaceCount--
	jCE.writeLine("</whileStatement>")
}

func (jCE *JackCompilationEngine) CompileIf() {
	jCE.writeLine("<ifStatement>")
	jCE.doubleSpaceCount++

	jCE.writeLine("<keyword> " + jCE.tokenizer.KeyWord() + " </keyword>")
	jCE.tokenizer.Advance()

	jCE.writeLine("<symbol> ( </symbol>")
	jCE.tokenizer.Advance()

	jCE.CompileExpression()

	jCE.writeLine("<symbol> ) </symbol>")
	jCE.tokenizer.Advance()

	jCE.writeLine("<symbol> { </symbol>")
	jCE.tokenizer.Advance()

	jCE.CompileStatements()

	jCE.writeLine("<symbol> } </symbol>")

	jCE.tokenizer.Advance()

	if jCE.tokenizer.TokenType() == analyzer.Keyword && jCE.tokenizer.KeyWord() == "else" {
		jCE.writeLine("<keyword> " + jCE.tokenizer.KeyWord() + " </keyword>")
		jCE.tokenizer.Advance()

		jCE.writeLine("<symbol> { </symbol>")
		jCE.tokenizer.Advance()

		jCE.CompileStatements()

		jCE.writeLine("<symbol> } </symbol>")

		jCE.tokenizer.Advance()
	}

	jCE.doubleSpaceCount--
	jCE.writeLine("</ifStatement>")
}

func (jCE *JackCompilationEngine) CompileLet() {
	jCE.writeLine("<letStatement>")
	jCE.doubleSpaceCount++

	jCE.compileKeywordOrIdentifiers()

	if jCE.tokenizer.TokenType() == analyzer.Symbol && jCE.tokenizer.Symbol() == '[' {
		jCE.writeLine("<symbol> [ </symbol>")
		jCE.tokenizer.Advance()

		jCE.CompileExpression()

		jCE.writeLine("<symbol> ] </symbol>")
		jCE.tokenizer.Advance()
	}
	jCE.writeLine("<symbol> = </symbol>")
	jCE.tokenizer.Advance()
	jCE.CompileExpression()

	jCE.writeLine("<symbol> ; </symbol>")
	jCE.tokenizer.Advance()

	jCE.doubleSpaceCount--
	jCE.writeLine("</letStatement>")
}

func (jCE *JackCompilationEngine) CompileExpression() {
	jCE.writeLine("<expression>")
	jCE.doubleSpaceCount++
	jCE.CompileTerm()

	for jCE.tokenizer.TokenType() == analyzer.Symbol && Operands[jCE.tokenizer.Symbol()] {
		jCE.writeLine("<symbol> " + html.EscapeString(string(jCE.tokenizer.Symbol())) + " </symbol>")
		jCE.tokenizer.Advance()
		jCE.CompileTerm()
	}

	jCE.doubleSpaceCount--
	jCE.writeLine("</expression>")
}

func (jCE *JackCompilationEngine) CompileTerm() {
	jCE.writeLine("<term>")
	jCE.doubleSpaceCount++

	if jCE.tokenizer.TokenType() == analyzer.Keyword {
		jCE.writeLine("<keyword> " + jCE.tokenizer.KeyWord() + " </keyword>")
	} else {
		jCE.writeLine("<identifier> " + jCE.tokenizer.Identifier() + " </identifier>")
	}
	jCE.tokenizer.Advance()

	jCE.doubleSpaceCount--
	jCE.writeLine("</term>")
}

func (jCE *JackCompilationEngine) isStatement() bool {
	return jCE.tokenizer.TokenType() == analyzer.Keyword && (jCE.tokenizer.KeyWord() == "let" ||
		jCE.tokenizer.KeyWord() == "if" ||
		jCE.tokenizer.KeyWord() == "while" ||
		jCE.tokenizer.KeyWord() == "do" ||
		jCE.tokenizer.KeyWord() == "return")
}

func (jCE *JackCompilationEngine) CompileVarDec() {
	jCE.writeLine("<varDec>")
	jCE.doubleSpaceCount++

	jCE.compileKeywordOrIdentifiers()

	for jCE.tokenizer.TokenType() == analyzer.Symbol && jCE.tokenizer.Symbol() == ',' {
		jCE.writeLine("<symbol> , </symbol>")
		jCE.tokenizer.Advance()
		jCE.compileKeywordOrIdentifiers()
	}
	jCE.writeLine("<symbol> ; </symbol>")
	jCE.tokenizer.Advance()

	jCE.doubleSpaceCount--
	jCE.writeLine("</varDec>")
}

func (jCE *JackCompilationEngine) CompileParameterList() {
	jCE.writeLine("<parameterList>")
	jCE.doubleSpaceCount++

	jCE.compileKeywordOrIdentifiers()
	for jCE.tokenizer.TokenType() == analyzer.Symbol && jCE.tokenizer.Symbol() == ',' {
		jCE.writeLine("<symbol> , </symbol>")
		jCE.tokenizer.Advance()
		jCE.compileKeywordOrIdentifiers()
	}

	jCE.doubleSpaceCount--
	jCE.writeLine("</parameterList>")
}

func (jCE *JackCompilationEngine) compileKeywordOrIdentifiers() {
	for jCE.tokenizer.TokenType() == analyzer.Keyword || jCE.tokenizer.TokenType() == analyzer.Identifier {
		if jCE.tokenizer.TokenType() == analyzer.Keyword {
			jCE.writeLine("<keyword> " + jCE.tokenizer.KeyWord() + " </keyword>")
		} else {
			jCE.writeLine("<identifier> " + jCE.tokenizer.Identifier() + " </identifier>")
		}
		jCE.tokenizer.Advance()
	}
}

func (jCE *JackCompilationEngine) isSubroutineDec() bool {
	return jCE.tokenizer.TokenType() == analyzer.Keyword &&
		(jCE.tokenizer.KeyWord() == "constructor" ||
			jCE.tokenizer.KeyWord() == "function" ||
			jCE.tokenizer.KeyWord() == "method")
}

func (jCE *JackCompilationEngine) CompileClassVarDec() {
	jCE.writeLine("<classVarDec>")
	jCE.doubleSpaceCount++

	jCE.compileKeywordOrIdentifiers()

	for jCE.tokenizer.TokenType() == analyzer.Symbol && jCE.tokenizer.Symbol() == ',' {
		jCE.writeLine("<symbol> , </symbol>")
		jCE.tokenizer.Advance()
		jCE.compileKeywordOrIdentifiers()
	}
	jCE.writeLine("<symbol> ; </symbol>")
	jCE.tokenizer.Advance()

	jCE.doubleSpaceCount--
	jCE.writeLine("</classVarDec>")
}

func (jCE *JackCompilationEngine) isClassVar() bool {
	return jCE.tokenizer.TokenType() == analyzer.Keyword && jCE.tokenizer.KeyWord() == "static" || jCE.tokenizer.KeyWord() == "field"
}

func (jCE *JackCompilationEngine) writeLine(line string) (int, error) {
	return jCE.sb.WriteString(spaces(jCE.doubleSpaceCount) + line + "\r\n")
}

func spaces(count int) string {
	var sb strings.Builder
	for i := 0; i < count; i++ {
		sb.WriteString(" ")
		sb.WriteString(" ")
	}
	return sb.String()
}
