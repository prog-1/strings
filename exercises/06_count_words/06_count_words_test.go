package main

import (
	"fmt"
	"testing"
)

func TestCountWords(t *testing.T) {
	pass := true
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
		{"string with numbers", "234 66 4", 3},
		{"mixed string", "11 abc ", 2},
	} {
		if got := countWords(tc.input); got != tc.want {
			fmt.Printf("%s\n\tFAIL: got %v, want %v\n", tc.name, got, tc.want)
			pass = false
		}
	}
	if pass {
		fmt.Println("ALL PASS")
	}
}
