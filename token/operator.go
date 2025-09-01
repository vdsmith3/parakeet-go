package token

func AdditionToken() *Token {
	return &Token{Type: ADDITION, Raw: "+"}
}

func SubtractionToken() *Token {
	return &Token{Type: SUBTRACTION, Raw: "-"}
}

func MultiplicationToken() *Token {
	return &Token{Type: MULTIPLICATION, Raw: "*"}
}

func DivisionToken() *Token {
	return &Token{Type: DIVISION, Raw: "/"}
}

func ModulusToken() *Token {
	return &Token{Type: MODULUS, Raw: "%"}
}

func ExponentToken() *Token {
	return &Token{Type: EXPONENT, Raw: "^"}
}

func AmpersandToken() *Token {
	return &Token{Type: AMPERSAND, Raw: "&"}
}

func PipeToken() *Token {
	return &Token{Type: PIPE, Raw: "|"}
}
