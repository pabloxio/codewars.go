package main_test

import (
	"testing"
	kata "github.com/pabloxio/codewars/6-kyu/human_readable_time"
)

func TestHumanReadableTime(t *testing.T) {
	tests := []struct{
		seconds int
		want string
	}{
		{0, "00:00:00"},
		{59, "00:00:59"},
		{60, "00:01:00"},
		{90, "00:01:30"},
		{3599, "00:59:59"},
		{3600, "01:00:00"},
		{45296, "12:34:56"},
		{86399, "23:59:59"},
		{86400, "24:00:00"},
		{359999, "99:59:59"},
	}

	for _, test := range tests {
		got := kata.HumanReadableTime(test.seconds)
		if got != test.want {
			t.Errorf("got %s, want %s", got, test.want)
		}
	}
}
