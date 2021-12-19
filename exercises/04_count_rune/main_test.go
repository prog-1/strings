package main

import "testing"

func TestCountRune(t *testing.T) {
	for _, tc := range []struct {
		input1 string
		input2 rune
		want   int
	}{
		{"", 98, 0},
		{"aaa", 97, 3},
		{"cake", 107, 1},
		{"sum", 115, 1},
		{"čūska", 269, 1},
	} {
		got := CountRune(tc.input1, tc.input2)
		if got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
