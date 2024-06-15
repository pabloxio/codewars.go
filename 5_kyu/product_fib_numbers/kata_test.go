package kata

import (
	"reflect"
	"testing"
)

func TestProductFib(t *testing.T) {
	tests := []struct {
		product uint64
		want    [3]uint64
	}{
		{4895, [3]uint64{55, 89, 1}},
		{5895, [3]uint64{89, 144, 0}},
		{74049690, [3]uint64{6765, 10946, 1}},
		{84049690, [3]uint64{10946, 17711, 0}},
	}

	for _, tt := range tests {
		got := ProductFib(tt.product)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v want %v", got, tt.want)
		}
	}
}
