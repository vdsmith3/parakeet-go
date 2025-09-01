package token

import "testing"

func TestIsWhitespace(t *testing.T) {
	tests := []struct {
		name string
		b    byte
		want bool
	}{
		{
			name: "space",
			b:    ' ',
			want: true,
		},
		{
			name: "tab",
			b:    '\t',
			want: true,
		},
		{
			name: "newline",
			b:    '\n',
			want: true,
		},
		{
			name: "carriage return",
			b:    '\r',
			want: true,
		},
		{
			name: "not whitespace",
			b:    'a',
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWhitespace(tt.b); got != tt.want {
				t.Errorf("IsWhitespace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWhitespaceToken(t *testing.T) {
	tests := []struct {
		name    string
		raw     byte
		want    *Token
		wantErr bool
	}{
		{
			name:    "space",
			raw:     ' ',
			want:    &Token{Type: WHITESPACE, Raw: " "},
			wantErr: false,
		},
		{
			name:    "tab",
			raw:     '\t',
			want:    &Token{Type: WHITESPACE, Raw: "\t"},
			wantErr: false,
		},
		{
			name:    "newline",
			raw:     '\n',
			want:    &Token{Type: WHITESPACE, Raw: "\n"},
			wantErr: false,
		},
		{
			name:    "carriage return",
			raw:     '\r',
			want:    &Token{Type: WHITESPACE, Raw: "\r"},
			wantErr: false,
		},
		{
			name:    "not whitespace",
			raw:     'a',
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WhitespaceToken(tt.raw)
			if got != nil && got.Type != tt.want.Type {
				t.Errorf("WhitespaceToken() = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("WhitespaceToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
