package main

import "testing"

func TestToLower(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"one letter", "a", "a"},
		{"one letter", "A", "a"},
		{"two letters", "ab", "ab"},
		{"two letters", "AB", "ab"},
		{"one word", "golang", "golang"},
		{"one word", "gOlAng", "golang"},
		{"one word", "GOLANG", "golang"},
		{"two words", "hOliDayS aFterNOoN", "holidays afternoon"},
		{"number", "123", "123"},
		{"mixed", " UtrO, veCHer! ", " utro, vecher! "},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToLower(tc.input); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"one letter", "a", "A"},
		{"one letter", "A", "A"},
		{"two letters", "ab", "AB"},
		{"two letters", "AB", "AB"},
		{"one word", "golang", "GOLANG"},
		{"one word", "gOlAng", "GOLANG"},
		{"one word", "GOLANG", "GOLANG"},
		{"two words", "hOliDayS aFterNOoN", "HOLIDAYS AFTERNOON"},
		{"number", "123", "123"},
		{"mixed", " UtrO, veCHer! ", " UTRO, VECHER! "},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToUpper(tc.input); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
