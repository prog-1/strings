package main

import "testing"

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		input2 byte
		want   int
	}{
		{"letter is not repeated", "mango", byte('l'), 0},
		{"letter is not repeated", "mango", byte('a'), 1},
		{"letter is repeated twice", "hello", byte('p'), 0},
		{"letter is repeated twice", "hello", byte('e'), 1},
		{"letter is repeated twice", "hello", byte('l'), 2},
		{"Latvian letters", "abčdē", 0xc4, 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountByte(tc.input, tc.input2); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
