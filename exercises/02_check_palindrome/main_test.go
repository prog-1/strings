package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  bool
	}{
		{"", true},
		{"abc", false},
		{"aaa", true},
		{"aba", true},
		{"123", false},
		{"242", true},
		{"a", true},
		{"ū", true},
		{"uūu", true},
		{"ab", false},
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
			if got := IsPalindrome(tc.input); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
