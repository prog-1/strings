package main

import "testing"

func TestCountByte(t *testing.T) {
	for _, tc := range []struct {
		input1 string
		input2 byte
		want   int
	}{
		{"hello", 108, 2},
		{"world", 111, 1},
		{"Привет", 109, 0},
		{"two", 99, 0},
		{"yigchfxn", 105, 1},
	} {
		got := CountByte(tc.input1, tc.input2)
		if got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
