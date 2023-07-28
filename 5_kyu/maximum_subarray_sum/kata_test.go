package kata

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, tt := range tests {
		got := MaximumSubarraySum(tt.input)
		if got != tt.want {
			t.Errorf("MaximumSubarraySum(%v), got %d, want %d", tt.input, got, tt.want)
		}
	}
}
