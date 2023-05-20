// https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go
package main

import "strings"

func ToCamelCase(s string) string {
	delimiter := "-"
	if strings.Contains(s, "_") {
		delimiter = "_"
	}

	words := strings.Split(s, delimiter)

	for i, word := range words[1:] {
		words[i+1] = strings.Title(word)
	}

	return strings.Join(words, "")
}
