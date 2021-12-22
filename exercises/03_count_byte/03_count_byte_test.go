package main

import "testing"

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		inputb byte
		want   int
	}{
		{"empty", "", 97, 0},
		{"one letter", "a", 97, 1},
		{"one word", "pizza", 97, 1},
		{"one word without needed byte", "pizza", 98, 0},
		{"one word with 2 same letters", "pizza", 122, 2},
		{"number", "12", 49, 1},
		{"number and word", "1 abc", 97, 1},
		{"2 words with same letters", "krak krek", 107, 4},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountByte(tc.input, tc.inputb); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
