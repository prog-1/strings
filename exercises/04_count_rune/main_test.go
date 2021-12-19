package main

import "testing"

func TestCountRune(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		input2 rune
		want   int
	}{
		{"letter is not repeated", "mango", 'l', 0},
		{"letter is not repeated", "mango", 'a', 1},
		{"letter is repeated twice", "hello", 'p', 0},
		{"letter is repeated twice", "hello", 'e', 1},
		{"letter is repeated twice", "hello", 'l', 2},
		{"Latvian letters", "abčdē", 'č', 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountRune(tc.input, tc.input2); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
