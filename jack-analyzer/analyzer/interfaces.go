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
	CompileClass() error
	CompileClassVarDec() error
	CompileSubroutineDec() error
	CompileParameterList() error
	CompileSubroutineBody() error
	CompileVarDec() error
	CompileStatements() error

	CompileLet() error
	CompileIf() error
	CompileWhile() error
	CompileDo() error
	CompileReturn() error

	CompileExpression() error
	CompileTerm() error
}
