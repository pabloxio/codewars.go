// https://www.codewars.com/kata/52fba66badcd10859f00097e/train/go
package main

func Disemvowel(comment string) string {
	var result []rune

	for _, v := range comment {
		switch v {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			continue
		default:
			result = append(result, v)
		}
	}

	return string(result)
}
