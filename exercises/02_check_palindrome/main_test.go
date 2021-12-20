package main

import "testing"

func TestIsPalindrom(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  bool
	}{
		{"", true},
		{"abc", false},
		{"aaa", true},
		{"123", false},
		{"242", true},
		{"a", true},
		{"ū", true},
		{"uūu", true},
		{"bb", true},
		{"33", true},
		{"dc", false},
		{"91", false},
		{"йй", true},
		{"йя", false},
		{"ātram slidas sadils martā", true},
		{"uūr", false},
		{"   abc   ", false},
	} {
		t.Run(tc.input, func(t *testing.T) {
			if got := IsPalindrom(tc.input); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
