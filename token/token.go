package token

type Token struct {
	Type TokenType
	Raw  string
}

type TokenType int

const (
	EOF = iota
)

func EOFToken() Token {
	return Token{Type: EOF, Raw: ""}
}
