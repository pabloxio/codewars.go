// https://www.codewars.com/kata/52774a314c2333f0a7000688/train/go
package main

import "strings"

func ValidParentheses(parens string) bool {
	if len(parens) == 0 {
		return true
	}

	if !strings.Contains(parens, "()") {
		return false
	}

	return ValidParentheses(strings.Join(strings.Split(parens, "()"), ""))
}
