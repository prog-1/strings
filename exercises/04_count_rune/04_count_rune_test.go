package main

import "testing"

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		name  string
		str   string
		srune rune
		want  int
	}{
		{"hello", "hello", 'l', 2},
		{"empty", "", 'l', 0},
		{"Latvian letters", "abčdē", 'č', 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountRune(tc.str, tc.srune); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
