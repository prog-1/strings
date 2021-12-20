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
		{"one word", "golang", "golang"},
		{"one word", "gOlAng", "golang"},
		{"one word", "GOLANG", "golang"},
		{"two words", "iron MAn", "iron man"},
		{"number", "234", "234"},
		{"mixed", "ČetveRg", "četverg"},
	} {
		got := ToLower(tc.input)
		if got != tc.want {
			t.Errorf("got = %q, want = %q", got, tc.want)
		}
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
		{"one word", "golang", "GOLANG"},
		{"one word", "gOlAng", "GOLANG"},
		{"one word", "GOLANG", "GOLANG"},
		{"two words", "iron MAn", "IRON MAN"},
		{"number", "234", "234"},
		{"mixed", "ČetveRg", "ČETVERG"},
	} {
		got := ToUpper(tc.input)
		if got != tc.want {
			t.Errorf("got = %q, want = %q", got, tc.want)
		}
	}
}
