package main

import "testing"

func TestGetMiddle(t *testing.T) {
	testcases := []struct {
		input, want string
	}{
		{"test", "es"},
		{"testing", "t"},
		{"middle", "dd"},
		{"A", "A"},
	}

	for _, tc := range testcases {
		t.Run(tc.input, func(t *testing.T) {
			got := GetMiddle(tc.input)

			if got != tc.want {
				t.Fatalf("GetMiddle(%q) = %q, want %q", tc.input, got, tc.want)
			}
		})
	}
}
