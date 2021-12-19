package main

import "testing"

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		input2 string
		want   int
	}{
		{"empty", "", "", 0},
		{"one letter", "a", "a", 0},
		{"number", "bananas", "ana", 2},
		{"word", "airplane", "a", 2},
		{"two words", "hello World hello", "ello", 1},
		{"word with unicode characters", "Hello    /t", " ", 4},
		{"symbols", "bababrdav1y2", "ba", 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Count(tc.input, tc.input2); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
