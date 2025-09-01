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
func Lex(input string) ([]*token.Token, error) {
	reader := strings.NewReader(input)
	return lex(reader)
}

func lex(reader *strings.Reader) ([]*token.Token, error) {
	nextToken, err := getNextToken(reader)
	if err != nil {
		return nil, err
	}

	if *nextToken == *token.EOFToken() {
		return []*token.Token{nextToken}, nil
	}

	followingTokens, err := lex(reader)
	if err != nil {
		return nil, err
	}

	if nextToken.Type == token.WHITESPACE {
		return followingTokens, nil
	}
	return append([]*token.Token{nextToken}, followingTokens...), nil
}

func getNextToken(reader *strings.Reader) (*token.Token, error) {
	if reader.Len() == 0 {
		return token.EOFToken(), nil
	}

	token, err := GetNumericToken(reader)
	if err == nil && token != nil {
		return token, nil
	}

	token, err = getWhitespaceToken(reader)
	if err == nil && token != nil {
		return token, nil
	}

	if err == nil {
		err = errors.New("could not lex token")
	}
	return nil, err
}

func peek(reader *strings.Reader) (byte, error) {
	peek, err := reader.ReadByte()
	if err != nil && err != io.EOF {
		return 0, err
	}
	reader.UnreadByte()
	return peek, nil
}

func isWhitespace(b byte) bool {
	return b == ' ' || b == '\t' || b == '\n' || b == '\r'
}

func getWhitespaceToken(reader *strings.Reader) (*token.Token, error) {
	peek, err := peek(reader)
	if err != nil {
		return nil, err
	}

	if isWhitespace(peek) {
		whitespaceByte, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		return token.WhitespaceToken(whitespaceByte), nil
	}

	return nil, nil
}
