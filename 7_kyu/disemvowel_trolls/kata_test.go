package main

import "testing"

func TestDisemvowel(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"This website is for losers LOL!", "Ths wbst s fr lsrs LL!"},
	}

	for _, test := range tests {
		if got := Disemvowel(test.input); got != test.want {
			t.Errorf("Disemvowel(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}
