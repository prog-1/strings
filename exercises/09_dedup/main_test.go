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

func TestDedup(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"one letter", "a", "a"},
		{"number", "120", "120"},
		{"number with duplicate", "120000", "120"},
		{"word", "airplaaaannne", "airplane"},
		{"two words", "heeelllllo worldddddd", "helo world"},
		{"word with spaces", "hello     ", "helo "},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Dedup(tc.input); !equal(got, tc.want) {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
