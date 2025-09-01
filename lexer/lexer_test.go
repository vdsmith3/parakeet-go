package lexer

import (
	"reflect"
	"testing"

	"github.com/vdsmith3/parakeet/token"
)

func TestLex(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []*token.Token
		wantErr bool
	}{
		{
			name:    "empty string",
			input:   "",
			want:    []*token.Token{},
			wantErr: false,
		},
		{
			name:    "single digit",
			input:   "1",
			want:    []*token.Token{{Type: token.INT, Raw: "1", Value: 1}},
			wantErr: false,
		},
		{
			name:  "multiple digits",
			input: "1 2 3  4",
			want: []*token.Token{
				{Type: token.INT, Raw: "1", Value: 1},
				{Type: token.INT, Raw: "2", Value: 2},
				{Type: token.INT, Raw: "3", Value: 3},
				{Type: token.INT, Raw: "4", Value: 4},
			},
			wantErr: false,
		},
		{
			name:  "multiple digits, variable whitespace",
			input: "1 2   3      4.152",
			want: []*token.Token{
				{Type: token.INT, Raw: "1", Value: 1},
				{Type: token.INT, Raw: "2", Value: 2},
				{Type: token.INT, Raw: "3", Value: 3},
				{Type: token.FLOAT, Raw: "4.152", Value: 4.152},
			},
			wantErr: false,
		},
		{
			name:  "operators",
			input: "1 + 2 - 3 * 4 / 5 % 6 ^ 7 & 8 | 9 && 10 || 11",
			want: []*token.Token{
				{Type: token.INT, Raw: "1", Value: 1},
				{Type: token.ADDITION, Raw: "+"},
				{Type: token.INT, Raw: "2", Value: 2},
				{Type: token.SUBTRACTION, Raw: "-"},
				{Type: token.INT, Raw: "3", Value: 3},
				{Type: token.MULTIPLICATION, Raw: "*"},
				{Type: token.INT, Raw: "4", Value: 4},
				{Type: token.DIVISION, Raw: "/"},
				{Type: token.INT, Raw: "5", Value: 5},
				{Type: token.MODULUS, Raw: "%"},
				{Type: token.INT, Raw: "6", Value: 6},
				{Type: token.EXPONENT, Raw: "^"},
				{Type: token.INT, Raw: "7", Value: 7},
				{Type: token.AMPERSAND, Raw: "&"},
				{Type: token.INT, Raw: "8", Value: 8},
				{Type: token.PIPE, Raw: "|"},
				{Type: token.INT, Raw: "9", Value: 9},
				{Type: token.AMPERSAND, Raw: "&"},
				{Type: token.AMPERSAND, Raw: "&"},
				{Type: token.INT, Raw: "10", Value: 10},
				{Type: token.PIPE, Raw: "|"},
				{Type: token.PIPE, Raw: "|"},
				{Type: token.INT, Raw: "11", Value: 11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Lex(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Lex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, append(tt.want, token.EOFToken())) {
				t.Errorf("Lex() = %v, want %v", got, tt.want)
			}
		})
	}
}
