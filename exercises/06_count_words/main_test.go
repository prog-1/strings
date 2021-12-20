package main

import (
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
		{"words plus digits", "1a 1 b1", 3},
		{"unicode", "привет, мир!", 2},
		{"mixed unicode", "čeлovek", 1},
		{"words plus punctuation", "cat, dog, fish.", 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountWords(tc.input); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
