package main

import "testing"

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		input2 byte
		want   int
	}{
		{"empty", "", 98, 0},
		{"one letter", "a", 97, 1},
		{"number", "1203", 50, 1},
		{"word", "airplane", 114, 1},
		{"two words", "hello World", 111, 2},
		{"word with unicode characters", "Hello    /t", 108, 2},
		{"symbols", "bababrdav1y2", 98, 3},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountByte(tc.input, tc.input2); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
