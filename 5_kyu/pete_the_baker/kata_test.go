package kata

import "testing"

func TestCakes(t *testing.T) {
	tests := []struct {
		recipe, available map[string]int
		want              int
	}{
		{
			map[string]int{"flour": 500, "sugar": 200, "eggs": 1},
			map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200},
			2,
		},
		{
			map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100},
			map[string]int{"sugar": 500, "flour": 2000, "milk": 2000},
			0,
		},
	}

	for _, tt := range tests {
		got := Cakes(tt.recipe, tt.available)

		if got != tt.want {
			t.Errorf("Cakes(%v, %v), got %d, want %d", tt.recipe, tt.available, got, tt.want)
		}
	}
}
