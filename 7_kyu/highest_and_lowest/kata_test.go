package main

import "testing"

func TestHighAndLow(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1 2 3 4 5", "5 1"},
		{"1 2 -3 4 5", "5 -3"},
		{"1 9 3 4 -5", "9 -5"},
		{"8 3 -5 42 -1 0 0 -9 4 7 4 -4", "42 -9"},
	}

	for _, test := range tests {
		if got := HighAndLow(test.input); got != test.want {
			t.Errorf("HighAndLow(%s) = %s, want %s", test.input, got, test.want)
		}
	}
}
