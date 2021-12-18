package main

import "testing"

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		inputb byte
		want   int
	}{
		{"empty", "", 101, 0},
		{"one letter", "h", 104, 1},
		{"number", "1203", 50, 1},
		{"mixed", "353gefbi5", 50, 0},
		{"mixed", "353gefbi5", 101, 1},
		{"one word", "genious", 110, 1},
		{"spces left", "    param", 97, 2},
		{"spaces right", "pam     ", 97, 1},
		{"two words", "tuc tuduc", 117, 3},
		{"mixed", "kc.cha!  u", 104, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountByte(tc.input, tc.inputb); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
