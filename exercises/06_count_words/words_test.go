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
		{"single word and spaces left", "   abc", 1},
		{"single word and spaces right", "abc   ", 1},
		{"single word and spaces everywhere", "   abc   ", 1},
		{"two words", "abc def", 2},
		{"two words and spaces everywhere", "   abc   def   ", 2},
		{"special letters", "čempionы syeē", 2},
	} {
		got := CountWords(tc.input)
		if got != tc.want {
			fmt.Printf("%s\n\tFAIL: got %v, want %v\n", tc.name, got, tc.want)
		}
	}
}
