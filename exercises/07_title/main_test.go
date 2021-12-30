package main

import (
	"testing"
)

func TestTitle(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"just one letter", "a", "A"},
		{"one word", "cHRISTmas", "Christmas"},
		{"one word in spaces", "   Andrej   ", "   Andrej   "},
		{"two words", "heLLO WORLD", "Hello World"},
		{"two words in spaces", "   hApPy  BiRtHdAy   ", "   Happy  Birthday   "},
		{"words with all letters lower case", "good morning", "Good Morning"},
		{"words with all letters upper case", "GOOD MORNING", "Good Morning"},
		{"numbers", "12334445678910 124", "12334445678910 124"},
		{"unicode", "привет, мир!", "Привет, Мир!"},
		{"mixed unicode", "čEЛoveK", "Čeлovek"},
		{"words plus punctuation", "Cat, dOg, fiSH.", "Cat, Dog, Fish."},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Title(tc.input); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
