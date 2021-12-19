package main

import "testing"

func TestIsPalidndrom(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  bool
	}{
		{"", true},
		{"abc", false},
		{"aaa", true},
		{"242", true},
		{"a", true},
		{"hello", false},
		{"ы", true},
		{"ufūi", false},
		{"ātram slidas sadils martā", true},
	} {
		t.Run(tc.input, func(t *testing.T) {
			if got := IsPalindrom(tc.input); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
