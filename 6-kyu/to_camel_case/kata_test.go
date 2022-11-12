package main_test

import (
	"testing"
	kata "github.com/pabloxio/codewars/6-kyu/to_camel_case"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct{
		s string
		want string
	}{
		{"the-stealth-warrior", "theStealthWarrior"},
		{"The_Stealth_Warrior", "TheStealthWarrior"},
	}

	for _, test := range tests {
		got := kata.ToCamelCase(test.s)
		if got != test.want {
			t.Errorf("got %s, want %s", got, test.want)
		}
	}
}
