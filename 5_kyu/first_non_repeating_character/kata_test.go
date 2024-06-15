package kata

import "testing"

func TestFirstNonRepeating(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"a", "a"},
		{"stress", "t"},
		{"moonmen", "e"},
		{"", ""},
		{"abba", ""},
		{"aa", ""},
		{"~><#~><", "#"},
		{"hello world, eh?", "w"},
		{"sTreSS", "T"},
		{"Go hang a salami, I'm a lasagna hog!", ","},
	}

	for _, tt := range tests {
		got := FirstNonRepeating(tt.input)

		if got != tt.want {
			t.Errorf("got %#v but want %#v", got, tt.want)
		}
	}
}
