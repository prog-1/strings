package main

import "testing"

func TestToLower(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"one letter is upper", "worLd", "world"},
		{"some letters are upper", "woRLd", "world"},
		{"all letters are upper", "WORLD", "world"},
	} {
		if got := ToLower(tc.input); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
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
		{"one letter is low", "WORLd", "WORLD"},
		{"some letters are low", "WoRld", "WORLD"},
		{"all letters are low", "world", "WORLD"},
	} {
		if got := ToUpper(tc.input); got != tc.want {
			t.Errorf("got = %v, want = %v", got, tc.want)
		}
	}
}
