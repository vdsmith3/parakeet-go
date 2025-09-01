package lexer

import (
	"io"
	"strings"

	"github.com/vdsmith3/parakeet/token"
)

func GetNumericToken(reader *strings.Reader) (*token.Token, error) {
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

		if detectFloat(nextByte) {
			return parseFloatToken(append(bytes, nextByte), reader)
		}
	}

	intToken, err := token.IntToken(bytes)
	if err != nil {
		return nil, err
	}
	return intToken, nil
}

func detectFloat(b byte) bool {
	return b == '.'
}

func parseFloatToken(bytes []byte, reader *strings.Reader) (*token.Token, error) {
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

	floatToken, err := token.FloatToken(bytes)
	if err != nil {
		return nil, err
	}
	return floatToken, nil
}
