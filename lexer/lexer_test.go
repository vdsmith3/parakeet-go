package lexer

import (
	"reflect"
	"testing"

	"github.com/vdsmith3/parakeet/token"
)

func TestLex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []token.Token
	}{
		{
			name:  "empty string",
			input: "",
			want:  []token.Token{token.EOFToken()},
		},
		// {
		// 	name:  "single character",
		// 	input: "a",
		// 	want:  []Token{EOF{}},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lex(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Lex() = %v, want %v", got, tt.want)
			}
		})
	}
}
