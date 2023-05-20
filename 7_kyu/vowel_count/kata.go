// https://www.codewars.com/kata/54ff3102c1bad923760001f3/train/go
package main

import "strings"

func GetCount(str string) (count int) {
	for _, v := range str {
		if strings.ContainsAny("aeiou", string(v)) {
			count++
		}
	}

	return
}
