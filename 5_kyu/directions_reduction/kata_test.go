package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsReasonable(t *testing.T) {
	tests := []struct {
		currentDir, nextDir string
		want                bool
	}{
		{"NORTH", "SOUTH", false},
		{"NORTH", "WEST", true},
		{"WEST", "EAST", false},
		{"WEST", "NORTH", true},
		{"EAST", "WEST", false},
		{"EAST", "SOUTH", true},
		{"SOUTH", "NORTH", false},
		{"SOUTH", "WEST", true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s going %s", tt.currentDir, tt.nextDir), func(t *testing.T) {
			got := isReasonable(tt.currentDir, tt.nextDir)
			if got != tt.want {
				t.Errorf("isReasonable(%q, %q), got %v, want %v", tt.currentDir, tt.nextDir, got, tt.want)
			}
		})
	}
}

func TestDirReduc(t *testing.T) {
	tests := []struct {
		input, want []string
	}{
		{[]string{"NORTH", "SOUTH"}, []string{}},
		{[]string{"NORTH", "SOUTH", "EAST"}, []string{"EAST"}},
		{[]string{"NORTH", "SOUTH", "EAST", "WEST"}, []string{}},
		{[]string{"SOUTH", "EAST", "WEST", "NORTH", "WEST"}, []string{"WEST"}},
		{[]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}, []string{"WEST"}},
		{[]string{"NORTH", "WEST", "SOUTH", "EAST"}, []string{"NORTH", "WEST", "SOUTH", "EAST"}},
		{[]string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}, []string{"NORTH"}},
	}

	for _, tt := range tests {
		got := DirReduc(tt.input)
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("DirReduc(%q) = got %v, want %v", tt.input, got, tt.want)
		}
	}
}
