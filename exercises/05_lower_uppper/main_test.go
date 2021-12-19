package main

import "testing"

func equal(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestToUpper(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"one letter", "a", "A"},
		{"number", "1203", "1203"},
		{"word", "airplane", "AIRPLANE"},
		{"two words", "heLlo WoRld", "HELLO WORLD"},
		{"word with spaces", "HEllo     ", "HELLO     "},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToUpper(tc.input); !equal(got, tc.want) {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
func TestToLower(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"one letter", "A", "a"},
		{"number", "1203", "1203"},
		{"word", "AIRPLANE", "airplane"},
		{"two words", "HELLO WORLD", "hello world"},
		{"word with spaces", "HELLO     ", "hello     "},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := ToLower(tc.input); !equal(got, tc.want) {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
