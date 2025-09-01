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
