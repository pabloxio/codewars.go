package kata

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		input, want []int
	}{
		{[]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}, []int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0}},
	}

	for _, tt := range tests {
		got := MoveZeros(tt.input)
		if reflect.DeepEqual(got, tt.input) {
			t.Errorf("MoveZeros(%v), got %v, want %v", tt.input, got, tt.want)
		}
	}
}
