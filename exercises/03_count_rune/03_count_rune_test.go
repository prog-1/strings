package main

import "testing"

func TestCountRune(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		inputr rune
		want   int
	}{
		{"empty", "", 'a', 0},
		{"one letter", "a", 'a', 1},
		{"one word", "pizza", 'a', 1},
		{"one word without needed rune", "pizza", 'l', 0},
		{"one word with 2 same letters", "pizza", 'z', 2},
		{"number", "12", '1', 1},
		{"number and word", "1 abc", 'b', 1},
		{"number and word", "1 abc", '1', 1},
		{"2 words with same letters", "krak krek", 'k', 4},
		{"unicode", "привет", 'р', 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountRune(tc.input, tc.inputr); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
