package main_test

import (
	"testing"

	kata "github.com/pabloxio/codewars/5_kyu/isbn_10_validation"
)

func TestValidISBN10(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"1112223339", true},
		{"111222333", false},
		{"1112223339X", false},
		{"1234554321", true},
		{"1234512345", false},
		{"048665088X", true},
		{"X123456788", false},
		{"048665088x", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := kata.ValidISBN10(tt.input); got != tt.want {
				t.Errorf("kata.ValidISBN10(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
