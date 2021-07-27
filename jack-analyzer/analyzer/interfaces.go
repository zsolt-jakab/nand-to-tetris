package analyzer

type Tokenizer interface {
	Advance()
	HasMoreTokens() bool
	TokenType() TokenType
	KeyWord() string
	Symbol() uint8
	Identifier() string
	IntVal() int
	StringVal() string
}

type CompilationEngine interface {
	CompileClass() string
	CompileClassVarDec()
	CompileSubroutineDec()
	CompileParameterList()
	CompileSubroutineBody()
	CompileVarDec()
	CompileStatements()

	CompileLet()
	CompileIf()
	CompileWhile()
	CompileDo()
	CompileReturn()

	CompileExpression()
	CompileTerm()
}
