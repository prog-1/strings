package main

import "testing"

func TestCountRune(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		inputr rune
		want   int
	}{
		{"empty", "", 101, 0},
		{"one letter", "h", 104, 1},
		{"number", "1203", 50, 1},
		{"mixed", "353gefbi5", 50, 0},
		{"mixed", "353gefbi5", 101, 1},
		{"one word", "genious", 110, 1},
		{"two words", "tuc tuduc", 117, 3},
		{"mixed", "kc.cha!  u", 104, 1},
		{"mixed", "āāлā111//....mmm,,ūč", 257, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountRune(tc.input, tc.inputr); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
