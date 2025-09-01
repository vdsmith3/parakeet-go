package token

import "testing"

func TestIntToken(t *testing.T) {
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
			want:    &Token{Type: INT, Raw: "1", Value: 1},
			wantErr: false,
		},
		{
			name:    "negative",
			raw:     []byte("-1"),
			want:    &Token{Type: INT, Raw: "-1", Value: -1},
			wantErr: false,
		},
		{
			name:    "multiple digits",
			raw:     []byte("1234"),
			want:    &Token{Type: INT, Raw: "1234", Value: 1234},
			wantErr: false,
		},
		{
			name:    "multiple digits, negative",
			raw:     []byte("-1234"),
			want:    &Token{Type: INT, Raw: "-1234", Value: -1234},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IntToken(tt.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Type != tt.want.Type {
				t.Errorf("IntToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
