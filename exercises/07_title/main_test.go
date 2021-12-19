package main

import "testing"

func equal(a, b []string) bool {
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

func TestTitle(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  []string
	}{
		{"empty", "", []string{}},
		{"one letter", "a", []string{"A"}},
		{"number", "1203", []string{"1203"}},
		{"word", "airplane", []string{"Airplane"}},
		{"two words", "heLlo WoRld", []string{"Hello", "World"}},
		{"word with spaces", "HEllo     ", []string{"Hello"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Title(tc.input); !equal(got, tc.want) {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
