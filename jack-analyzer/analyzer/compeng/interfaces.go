package compeng

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
