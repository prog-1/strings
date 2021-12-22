package main

import "testing"

func TestCount(t *testing.T) {
	for _, tc := range []struct {
		name   string
		input  string
		inputs string
		want   int
	}{
		{"empty", "", "", 0},
		{"one letter", "a", "a", 0},
		{"word", "bananas", "ana", 2},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Count(tc.input); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
