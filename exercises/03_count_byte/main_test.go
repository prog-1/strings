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
		{"empty", "", 0, 0}, // {"empty", "", 'e', 0},
		{"one letter", "h", 'h', 1},
		{"two letters", "ha", 'a', 1},
		{"two-digit number", "13", '2', 0},
		{"three-digit number", "123", '2', 1},
		{"four-digit number", "2123", '2', 2},
		{"one word", "generous", 'n', 1},
		{"two words", "tuc tuduc", 'u', 3},
		{"spaces left", "    tadam", 'a', 2},
		{"spaces right", "tam     ", 'a', 1},
		{"mixed", "353gefbi5", '2', 0},
		{"mixed", "353GEf,bi5.", 'E', 1},
		{"mixed", "tuk.tuk!  and?", 'u', 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountByte(tc.input, tc.inputbyte); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}

		})
	}
}
