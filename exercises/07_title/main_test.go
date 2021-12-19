package main

import "testing"

func TestTitle(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  string
	}{
		{"", ""},
		{"hello", "Hello"},
		{"WORLD", "World"},
		{"пРИвЕт", "Привет"},
		{"čūSkA", "Čūska"},
		{"010111", "010111"},
	} {
		got := Title(tc.input)
		if got != tc.want {
			t.Errorf("got = %q, want = %q", got, tc.want)
		}
	}
}
