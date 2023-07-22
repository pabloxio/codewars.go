package main

import "testing"

func TestRot13(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"EBG13 rknzcyr.", "ROT13 example."},
		{"This is my first ROT13 excercise!", "Guvf vf zl svefg EBG13 rkprepvfr!"},
	}

	for _, tt := range tests {
		got := Rot13(tt.input)
		if got != tt.want {
			t.Errorf("Rot13(%q), got %q, want %q", tt.input, got, tt.want)
		}
	}
}
