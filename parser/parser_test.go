package parser

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "empty string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Parse("")
			if got != nil {
				t.Errorf("Parse() = %v, want nil", got)
			}
		})
	}
}
