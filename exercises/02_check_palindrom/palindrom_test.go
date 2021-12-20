package main

import "testing"

func TestIsPalindrom(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", true},
		{"one word", "abc", false},
		{"palindrom letters", "aaa", true},
		{"palindrom numbers", "242", true},
		{"one letter", "a", true},
		{"special letter", "ū", true},
		{"spaces", "   abc  ", false},
		{"mixed palinrom", "uūu", true},
		{"mixed plaindrom again", "ātram slidas sadils martā", true},
		{"mixed but not a palindrom", "čeлovek", false},
		{"spaces palindrom", "   aba   ", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsPalindrom(tc.input); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
