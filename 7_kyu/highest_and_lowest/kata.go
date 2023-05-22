// https://www.codewars.com/kata/554b4ac871d6813a03000035/train/go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	var numbers []int
	for _, n := range strings.Split(in, " ") {
		number, _ := strconv.Atoi(n)
		numbers = append(numbers, number)
	}

	high, low := numbers[0], numbers[0]

	for _, number := range numbers {
		if number > high {
			high = number
		}
		if number < low {
			low = number
		}
	}

	return fmt.Sprintf("%d %d", high, low)
}
