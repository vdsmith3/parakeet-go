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
			input: "1 2   3      4  ",
			want: []*token.Token{
				{Type: token.INT, Raw: "1", Value: 1},
				{Type: token.INT, Raw: "2", Value: 2},
				{Type: token.INT, Raw: "3", Value: 3},
				{Type: token.INT, Raw: "4", Value: 4},
			},
			wantErr: false,
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
