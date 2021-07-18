package analyzer

var Symbols = map[uint8]bool{
	'{': true,
	'}': true,
	'(': true,
	')': true,
	'[': true,
	']': true,
	'.': true,
	',': true,
	';': true,
	'+': true,
	'-': true,
	'*': true,
	'/': true,
	'&': true,
	'|': true,
	'<': true,
	'>': true,
	'=': true,
	'~': true,
}

var Keywords = map[string]bool{
	"class":       true,
	"constructor": true,
	"function":    true,
	"method":      true,
	"field":       true,
	"static":      true,
	"var":         true,
	"int":         true,
	"char":        true,
	"boolean":     true,
	"void":        true,
	"true":        true,
	"false":       true,
	"null":        true,
	"this":        true,
	"let":         true,
	"do":          true,
	"if":          true,
	"else":        true,
	"while":       true,
	"return":      true,
}

type TokenType int

const (
	Keyword TokenType = iota
	Symbol
	Identifier
	IntConst
	StringConst
	UnknownToken
)

func (tt TokenType) String() string {
	return []string{"keyword", "symbol", "identifier", "integerConstant", "stringConstant", "unknownToken"}[tt]
}
