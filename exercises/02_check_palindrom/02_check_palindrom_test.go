package main

import "testing"

func TestIsPalidndrom(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  bool
	}{
		{"empty", "", true},
		{"palidrom", "abccba", true},
		{"palidrom", "abcba", true},
		{"not palidrom", "abcdef", false},
		{"not palidrom", "abcabc", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsPalidrom(tc.input); got != tc.want {
				t.Errorf("isPalidndgot(%v) = %v, want = %v", tc.input, got, tc.want)
			}
		})
	}
}
