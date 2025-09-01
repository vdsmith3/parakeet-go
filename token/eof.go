package token

func (t *Token) IsEOF(token *Token) bool {
	return t.Type == EOF
}

func EOFToken() *Token {
	return &Token{Type: EOF, Raw: ""}
}
