// https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c/train/go
package kata

func MaximumSubarraySum(numbers []int) int {
	var sum int

	for i := 0; i < len(numbers)-1; i++ {
		tmpSum := numbers[i]

		for j := i + 1; j < len(numbers); j++ {
			tmpSum += numbers[j]
			if tmpSum > sum {
				sum = tmpSum
			}
		}
	}

	return sum
}
