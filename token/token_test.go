package token

import "testing"

// if it's ever something else EOF is borked
func TestEOFToken(t *testing.T) {
	got := EOFToken()
	want := Token{Type: EOF, Raw: ""}
	if got != want {
		t.Errorf("EOFToken() = %v, want %v", got, want)
	}
}
