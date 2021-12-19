package main

import "testing"

func TestCount(t *testing.T) {
	for _, tc := range []struct {
		input string
		subs  string
		want  int
	}{
		{"", "", 0},
		{"Hello", "", 0},
		{"", "ll", 0},
		{"world", "or", 1},
		{"first", "first", 1},
		{"машина", "шина", 1},
		{"bananas", "ana", 2},
	} {
		got := Count(tc.input, tc.subs)
		if got != tc.want {
			t.Errorf("got = %q, want = %v", got, tc.want)
		}
	}
}
