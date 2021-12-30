package main

import "testing"

func TestToLower(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"one lower case letter", "a", "a"},
		{"one upper case letter", "A", "a"},
		{"two lower case letters", "ab", "ab"},
		{"two upper case letters", "AB", "ab"},
		{"one word with all letters lower case", "hello", "hello"},
		{"one word with letters both lower case and upper case", "hEllO", "hello"},
		{"one word with all letters upper case", "HELLO", "hello"},
		{"two words with letters both lower case and upper case", "hOliDayS aFterNOoN", "holidays afternoon"},
		{"number", "123", "123"},
		{"mixed", " UtrO, veCHer! ", " utro, vecher! "},
		{"unicode", "Привет, миР!", "Привет, миР!"},           // if there is a non-English letter in the input, the program will just return it unchanged
		{"mixed unicode", "Priвет,  Дryг.", "priвет,  Дryг."}, // "y" in "дryг" is an English letter

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
		{"one lower case letter", "a", "A"},
		{"one upper case letter", "A", "A"},
		{"two lower case letters", "ab", "AB"},
		{"two upper case letters", "AB", "AB"},
		{"one word with all letters lower case", "hello", "HELLO"},
		{"one word with letters both lower case and upper case", "hEllO", "HELLO"},
		{"one word with all letters upper case", "HELLO", "HELLO"},
		{"two words with letters both lower case and upper case", "hOliDayS aFterNOoN", "HOLIDAYS AFTERNOON"},
		{"number", "123", "123"},
		{"mixed", " UtrO, veCHer! ", " UTRO, VECHER! "},
		{"unicode", "Привет, миР!", "Привет, миР!"},           // if there is a non-English letter in the input, the program will just return it unchanged
		{"mixed unicode", "Priвет,  Дryг.", "PRIвет,  ДRYг."}, // "y" in "дryг" is an English letter
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToUpper(tc.input); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
