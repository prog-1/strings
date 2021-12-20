package main

import "testing"

func TestTitle(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"symbols", " *-+!,128", " *-+!,.128"},
		{"one letter", "a", "A"},
		{"one letter", "A", "A"},
		{"single word", "abc", "Abc"},
		{"single word", "ABC", "Abc"},
		{"single word", "aBc", "Abc"},
		{"single word and spaces left", "   abc", "   Abc"},
		{"two words", "abc def", "Abc Def"},
		{"two words and spaces", "   abc   def", "   Abc   Def"},
		{"special letters", "čempions ffa", "Čempions Ffa"},
	} {
		got := Title(tc.input)
		if got != tc.want {
			t.Errorf("got = %q, want = %q", got, tc.want)
		}
	}
}
