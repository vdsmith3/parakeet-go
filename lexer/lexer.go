package lexer

import "github.com/vdsmith3/parakeet/token"

// Lexes the input string into a slice of tokens
// Outputs tokens in reverse order
//
// input  - The input string to lex
// output - A slice of tokens
func Lex(input string) []token.Token {
	nextToken, remainingInput := getNextToken(input)
	if nextToken.Type == token.EOF {
		return []token.Token{nextToken}
	}
	return append(Lex(remainingInput), nextToken)
}

// Lexes the next token in the input string
//
// input - The input string to lex
// output:
//   - The next token in the input string
//   - The remaining input string after the token
func getNextToken(input string) (token.Token, string) {
	if len(input) == 0 {
		return token.EOFToken(), ""
	}

	// TODO: Implement more tokens
	return token.Token{}, input
}
