package main

import "testing"

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		name  string
		str   string
		sbyte byte
		want  int
	}{
		{"hello", "hello", 0x6c, 2},
		{"empty", "", 0x6c, 0},
		{"Latvian letters", "abčdē", 0xc4, 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountByte(tc.str, tc.sbyte); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
