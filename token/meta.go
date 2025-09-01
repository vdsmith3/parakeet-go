package token

func (t *Token) IsEOF(token *Token) bool {
	return t.Type == EOF
}

func EOFToken() *Token {
	return &Token{Type: EOF, Raw: ""}
}

func WhitespaceToken(raw byte) *Token {
	return &Token{Type: WHITESPACE, Raw: string(raw)}
}
