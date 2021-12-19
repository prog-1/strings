package main

import "testing"

func TestToLower(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  string
	}{
		{"", ""},
		{"HeLlo", "hello"},
		{"SUrFinG", "surfing"},
		{"ПриВеТ", "привет"},
	} {
		got := ToLower(tc.input)
		if got != tc.want {
			t.Errorf("got = %q, want = %q", got, tc.want)
		}
	}
}

func TestToUpper(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  string
	}{
		{"", ""},
		{"HeLlo", "HELLO"},
		{"SUrFinG", "SURFING"},
		{"ПриВеТ", "ПРИВЕТ"},
	} {
		got := ToUpper(tc.input)
		if got != tc.want {
			t.Errorf("got = %q, want = %q", got, tc.want)
		}
	}
}
