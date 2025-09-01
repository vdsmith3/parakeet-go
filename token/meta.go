package token

import "errors"

func EOFToken() *Token {
	return &Token{Type: EOF, Raw: ""}
}

func IsWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func WhitespaceToken(raw byte) (*Token, error) {
	if !IsWhitespace(raw) {
		return nil, errors.New("not a whitespace token")
	}
	return &Token{Type: WHITESPACE, Raw: string(raw)}, nil
}
