package kata

import (
	"testing"
)

func TestRGB(t *testing.T) {
	tests := []struct {
		r, g, b int
		want    string
	}{
		{255, 255, 255, "FFFFFF"},
		{255, 255, 300, "FFFFFF"},
		{0, 0, 0, "000000"},
		{148, 0, 211, "9400D3"},
		{0, 0, 0, "000000"},
		{1, 2, 3, "010203"},
		{254, 253, 252, "FEFDFC"},
		{-20, 275, 125, "00FF7D"},
	}

	for _, tt := range tests {
		got := RGB(tt.r, tt.g, tt.b)

		if got != tt.want {
			t.Errorf("RGB(%d, %d, %d), got %s, want %s", tt.r, tt.g, tt.b, got, tt.want)
		}
	}
}
