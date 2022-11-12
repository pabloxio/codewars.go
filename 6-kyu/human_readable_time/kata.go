// https://www.codewars.com/kata/52685f7382004e774f0001f7/train/go
package main

import "fmt"

func stripTime(t int) (int, int) {
	return t/60, t%60
}

func HumanReadableTime(seconds int) string {
	remaining, seconds := stripTime(seconds)
	hours, minutes := stripTime(remaining)
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
