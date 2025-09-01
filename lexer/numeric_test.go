package lexer

import (
	"strings"
	"testing"

	"github.com/vdsmith3/parakeet/token"
)

func TestGetNumericToken(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *token.Token
		wantErr bool
	}{
		{
			name:    "empty string",
			input:   "",
			want:    nil,
			wantErr: false,
		},
		{
			name:    "1",
			input:   "1",
			want:    &token.Token{Type: token.INT, Raw: "1", Value: 1},
			wantErr: false,
		},
		{
			name:    "negative",
			input:   "-1",
			want:    &token.Token{Type: token.INT, Raw: "-1", Value: -1},
			wantErr: false,
		},
		{
			name:    "multiple digits",
			input:   "1234",
			want:    &token.Token{Type: token.INT, Raw: "1234", Value: 1234},
			wantErr: false,
		},
		{
			name:    "multiple digits, negative",
			input:   "-1234",
			want:    &token.Token{Type: token.INT, Raw: "-1234", Value: -1234},
			wantErr: false,
		},
		{
			name:    "float",
			input:   "1.234",
			want:    &token.Token{Type: token.FLOAT, Raw: "1.234", Value: 1.234},
			wantErr: false,
		},
		{
			name:    "float, negative",
			input:   "-1.234",
			want:    &token.Token{Type: token.FLOAT, Raw: "-1.234", Value: -1.234},
			wantErr: false,
		},
		{
			name:    "float, less than 1",
			input:   "0.1234",
			want:    &token.Token{Type: token.FLOAT, Raw: "0.1234", Value: 0.1234},
			wantErr: false,
		},
		{
			name:    "invalid float",
			input:   "1.234.567",
			want:    &token.Token{Type: token.FLOAT, Raw: "1.234", Value: 1.234},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			got, err := GetNumericToken(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumericToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Type != tt.want.Type {
				t.Errorf("GetNumericToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
