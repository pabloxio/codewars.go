// https://www.codewars.com/kata/52223df9e8f98c7aa7000062/train/go
package main

import "strings"

func Rot13(msg string) string {
	return strings.Map(func(r rune) rune {
		// Solution based on Wikipedia description
		// https://en.wikipedia.org/wiki/ROT13
		switch {
		case (r >= 'A' && r <= 'M') || (r >= 'a' && r <= 'm'):
			r += 13
		case (r >= 'N' && r <= 'Z') || (r >= 'n' && r <= 'z'):
			r -= 13
		}

		return r
	}, msg)
}
