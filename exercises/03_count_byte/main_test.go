package main

import (
	"testing"
)

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		name      string
		input     string
		inputbyte byte
		want      int
	}{
		{"empty", "", 101, 0},
		{"one letter", "h", 104, 1},
		{"two letters", "ha", 97, 1},
		{"two-digit number", "13", 50, 0},
		{"three-digit number", "123", 50, 1},
		{"three-digit number", "2123", 50, 2},
		{"one word", "generous", 110, 1},
		{"two words", "tuc tuduc", 117, 3},
		{"spaces left", "    tadam", 97, 2},
		{"spaces right", "tam     ", 97, 1},
		{"mixed", "353gefbi5", 50, 0},
		{"mixed", "353gef,bi5.", 101, 1},
		{"mixed", "tuk.tuk!  and?", 117, 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountByte(tc.input, tc.inputbyte); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}

		})
	}
}
