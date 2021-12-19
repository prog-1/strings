package main

import "testing"

func TestDedup(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  string
	}{
		{"", ""},
		{"aaaa", "a"},
		{"000111110001111", "0101"},
		{"hello", "helo"},
		{"hello wworllddddd", "helo world"},
		{"ппприввет", "привет"},
	} {
		got := Dedup(tc.input)
		if got != tc.want {
			t.Errorf("got = %q, want = %q", got, tc.want)
		}
	}
}
