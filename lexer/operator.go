package lexer

import (
	"strings"

	"github.com/vdsmith3/parakeet/token"
)

func getOperatorToken(reader *strings.Reader) (*token.Token, error) {
	peek, err := peek(reader)
	if err != nil {
		return nil, err
	}

	if isOperatorByte(peek) {
		return parseOperatorToken(reader)
	}
	return nil, nil // not an operand
}

func isOperatorByte(b byte) bool {
	return b == '+' || b == '-' || b == '*' || b == '/' || b == '%' || b == '^'
}

func parseOperatorToken(reader *strings.Reader) (*token.Token, error) {
	nextByte, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	operatorMap := map[byte]*token.Token{
		'+': token.AdditionToken(),
		'-': token.SubtractionToken(),
		'*': token.MultiplicationToken(),
		'/': token.DivisionToken(),
		'%': token.ModulusToken(),
		'^': token.ExponentToken(),
		'&': token.AmpersandToken(),
		'|': token.PipeToken(),
	}

	if operator, ok := operatorMap[nextByte]; ok {
		return operator, nil
	}
	return nil, nil
}
