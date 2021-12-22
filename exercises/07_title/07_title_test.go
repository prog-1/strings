package main

import "testing"

func TestTitle(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"one letter without upper", "a", "A"},
		{"one letter with upper(capital)", "A", "A"},
		{"two words", "abc def", "Abc Def"},
		{"two words (correct input)", "Abc Def", "Abc Def"},
		{"numbers", "123 44 5", "123 44 5"},
		{"one word with one capital letter (not start of string)", "aBc", "Abc"},
		{"one word with one capital letter (not start of string)", "abC", "AbC"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Title(tc.input); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
