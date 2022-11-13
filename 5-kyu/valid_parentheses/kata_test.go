package main_test

import (
	"testing"

	kata "github.com/pabloxio/codewars/5-kyu/valid_parentheses"
)


func TestValidParentheses(t *testing.T) {
	tests := []struct{
		parns string
		want bool
	}{
		{"()", true},
		{")(()))", false},
		{"(", false},
		{"(())((()())())", true},
		{"())(", false},
	}

	for _, test := range tests {
		got := kata.ValidParentheses(test.parns)

		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
