// https://www.codewars.com/kata/52bc74d4ac05d0945d00054e/train/go
package kata

import "strings"

func FirstNonRepeating(s string) string {
	for _, char := range s {
		if strings.Count(strings.ToLower(s), strings.ToLower(string(char))) == 1 {
			return string(char)
		}
	}

	return ""
}
