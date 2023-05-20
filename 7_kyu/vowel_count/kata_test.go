package main_test

import (
	"testing"

	kata "github.com/pabloxio/codewars/7_kyu/vowel_count"
)

func TestGetCount(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"abracadabra", 5},
		{"", 0},
		{"pear tree", 4},
		{"o a kak ushakov lil vo kashu kakao", 13},
		{"my pyx", 0},
		{"o", 1},
	}

	for _, test := range tests {
		if got := kata.GetCount(test.input); got != test.want {
			t.Errorf("GetCount(%q) = %d, want %d", test.input, got, test.want)
		}
	}
}
