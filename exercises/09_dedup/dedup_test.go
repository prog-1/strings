package main

import "testing"

func TestDedup(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"one letter", "a", "a"},
		{"repeating letters", "bbbbbbeeebbbbbbeeee", "bebe"},
		{"repeating numbers", "22888000", "280"},
		{"two words", "merry christmaas", "mery christmas"},
		{"special letters", "čččelovekkk", "čelovek"},
		{"spaces", "   ", " "},
		{"spaces and word", "  puf", " puf"},
		{"spaces and word", "puf  ", "puf "},
		{"spaces and word", "  puf  ", " puf "},
	} {
		got := Dedup(tc.input)
		if got != tc.want {
			t.Errorf("got = %q, want = %q", got, tc.want)
		}
	}
}
