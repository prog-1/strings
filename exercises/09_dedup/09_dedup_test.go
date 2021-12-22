package main

import "testing"

func TestDedup(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"one letter without duplicate", "a", "a"},
		{"one word without duplicate", "abc", "abc"},
		{"one word with duplicate", "piiiizzzzza", "piza"},
		{"number without", "1", "1"},
		{"number with duplicate", "000111110001111", "0101"},
		{"two words with duplicate", "piiizza tomaaato", "piza tomato"},
		{"unicode", "приввееет", "привет"},
		{"spaces", "   ", " "},
		{"left spaces and word", "ab   ", "ab "},
		{"right spaces and world", "   ab", " ab"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Dedup(tc.input); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
