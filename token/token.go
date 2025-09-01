package token

type Token struct {
	Type  TokenType
	Raw   string
	Value interface{}
}

type TokenType int

const (
	EOF = iota
	WHITESPACE
	INT
	FLOAT
)
