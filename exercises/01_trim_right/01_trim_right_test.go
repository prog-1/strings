package main

import "testing"

func TestTrimRight(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"nothing to trim", "abc", "abc"},
		{"trim set", " \n\r\t", ""},
		{"double trim set", "  \n\n\r\r\t\t", ""},
		{"mixed trim set", " \t \n \t \r \t \n", ""},
		{"nothing to trim", "\t  \n \r    abc", "\t  \n \r    abc"},
		{"keeps tail", "abc\t  \n \r    ", "abc"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := TrimRight(tc.input); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
