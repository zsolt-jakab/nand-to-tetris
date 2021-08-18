package tokenizer

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
