// https://www.codewars.com/kata/52597aa56021e91c93000cb0/trajn/go

package kata

func MoveZeros(arr []int) []int {
	output := make([]int, len(arr))
	var index int

	for _, value := range arr {
		if value != 0 {
			output[index] = value
			index++
		}
	}

	return output
}
