// https://www.codewars.com/kata/56747fd5cb988479af000028/train/go
package main

func GetMiddle(s string) string {
	even := len(s)%2 == 0

	middle := len(s) / 2
	if even {
		return s[middle-1 : middle+1]
	}

	return s[middle : middle+1]
}
