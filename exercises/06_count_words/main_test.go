package main

import (
	"fmt"
	"testing"
)

func TestCountWords(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"separators", " *-+!,.\n\r\t", 0},
		{"single word", "abc", 1},
		{"single word in spaces", "   abc   ", 1},
		{"two words", "abc def", 2},
		{"two words in spaces", "   abc   def   ", 2},
		{"unicode", "привет, мир!", 2},
	} {
		got := CountWords(tc.input)
		if got != tc.want {
			fmt.Printf("%s\n\tFAIL: got %v, want %v\n", tc.name, got, tc.want)
		}
	}
}
