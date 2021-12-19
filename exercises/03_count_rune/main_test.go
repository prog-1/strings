package main

import "testing"

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		input2 rune
		want   int
	}{
		{"empty", "", 'l', 0},
		{"one letter", "a", 'a', 1},
		{"number", "1203", '1', 1},
		{"word", "airplane", 'a', 2},
		{"two words", "hello World", 'l', 3},
		{"word with unicode characters", "Hello    /t", 'H', 1},
		{"symbols", "bababrdav1y2", 'a', 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountRune(tc.input, tc.input2); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
