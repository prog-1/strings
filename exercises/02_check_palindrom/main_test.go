package main

import "testing"

func TestIsPalidndrom(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", true},
		{"palindrom", "a", true},
		{"palindrom", "madam", true},
		{"palinrom with unicode characters", " madam ", true},
		{"not palidrom ", "abc", false},
		{"not palidrom with unicode characters", "abc    ", false},
	} {

		t.Run(tc.name, func(t *testing.T) {
			if got := IsPalindrome(tc.input); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
