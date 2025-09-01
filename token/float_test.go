package token

import "testing"

func TestFloatToken(t *testing.T) {
	tests := []struct {
		name    string
		raw     []byte
		want    *Token
		wantErr bool
	}{
		{
			name:    "empty string",
			raw:     []byte(""),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "1",
			raw:     []byte("1"),
			want:    &Token{Type: FLOAT, Raw: "1", Value: 1.0},
			wantErr: false,
		},
		{
			name:    "1.234",
			raw:     []byte("1.234"),
			want:    &Token{Type: FLOAT, Raw: "1.234", Value: 1.234},
			wantErr: false,
		},
		{
			name:    "1.234.567",
			raw:     []byte("1.234.567"),
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FloatToken(tt.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("FloatToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Type != tt.want.Type {
				t.Errorf("FloatToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
