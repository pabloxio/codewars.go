// https://www.codewars.com/kata/51fc12de24a9d8cb0e000001/train/go
package main

import (
	"regexp"
	"strconv"
	"strings"
)

func ValidISBN10(isbn string) bool {
	re, _ := regexp.Compile(`(?i)^\d{9}[0-9X]$`)
	if !re.Match([]byte(isbn)) {
		return false
	}

	var sum, num int

	for i, v := range isbn {
		if strings.ToUpper(string(v)) == "X" {
			num = 10
		} else {
			num, _ = strconv.Atoi(string(v))
		}

		sum += num * (i + 1)
	}

	return sum%11 == 0
}
