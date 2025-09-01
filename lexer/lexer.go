package lexer

import (
	"errors"
	"io"
	"strings"

	"github.com/vdsmith3/parakeet/token"
)

// Lexes the input string into a slice of tokens
// Outputs tokens in reverse order
//
// input  - The input string to lex
// output - A slice of tokens
func Lex(input string) ([]token.Token, error) {
	reader := strings.NewReader(input)
	return lex(reader)
}

func lex(reader *strings.Reader) ([]token.Token, error) {
	nextToken, err := getNextToken(reader)
	if err != nil {
		return nil, err
	}

	if *nextToken == *token.EOFToken() {
		return []token.Token{*nextToken}, nil
	}

	followingTokens, err := lex(reader)
	if err != nil {
		return nil, err
	}
	return append([]token.Token{*nextToken}, followingTokens...), nil
}

func peek(reader *strings.Reader) (byte, error) {
	peek, err := reader.ReadByte()
	if err != nil {
		return 0, err
	}
	reader.UnreadByte()
	return peek, nil
}

func getNextToken(reader *strings.Reader) (*token.Token, error) {
	if reader.Len() == 0 {
		return token.EOFToken(), nil
	}

	token, err := getNumericToken(reader)
	if err != nil {
		return nil, err
	}

	if token != nil {
		return token, nil
	}

	return nil, errors.New("could not lex token")
}

func getNumericToken(reader *strings.Reader) (*token.Token, error) {
	peek, err := peek(reader)
	if err != nil {
		return nil, err
	}

	if isNumericByte(peek) {
		return parseNumericalToken(reader)
	}
	return nil, nil // not a number
}

func isNumericByte(b byte) bool {
	return b >= '0' && b <= '9'
}

func parseNumericalToken(reader *strings.Reader) (*token.Token, error) {
	bytes := make([]byte, 0)
	nextByte, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	for isNumericByte(nextByte) {
		bytes = append(bytes, nextByte)
		nextByte, err = reader.ReadByte()
		if err != nil && err != io.EOF {
			return nil, err
		}
	}

	// unread the last byte that is not a digit
	// reader.UnreadByte()
	intToken, err := token.IntToken(bytes)
	if err != nil {
		return nil, err
	}
	return intToken, nil
}
