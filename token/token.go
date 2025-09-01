package token

type Token struct {
	Type  TokenType
	Raw   string
	Value interface{}
}

type TokenType int

const (
	// meta
	EOF = iota
	WHITESPACE

	// numeric
	INT
	FLOAT

	// operator
	ADDITION
	SUBTRACTION
	MULTIPLICATION
	DIVISION
	MODULUS
	EXPONENT
	AMPERSAND
	PIPE
)
