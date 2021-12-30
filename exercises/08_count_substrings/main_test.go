package main

import (
	"testing"
)

func TestCount(t *testing.T) {
	for _, tc := range []struct {
		name     string
		input    string
		inputsub string
		want     int
	}{
		{"empty", "", "", 0},
		{"one word without substring existing in it", "abc", "de", 0},
		{"one word with substring existing 1 time", "ban", "an", 1},
		{"one word with the same substring as the word", "apple juice", "apple juice", 1},
		{"one word with substring existing 2 times", "bananas", "ana", 2},
		{"two words without substring existing in it", "hello world", "heLLo", 0},
		{"two words with substring existing 1 time", "old friend", "fri", 1},
		{"two words with substring existing 3 times", "merry christmas", "r", 3},
		{"three words with substring existing 2 times", "Who are you and what are you doing?", "are you", 2},
		{"numbers", "123344456789!", "34445", 1},
		{"unicode", "Привет, Мир!", "Прив", 1},
		{"mixed unicode", "МОJ stAрый  ДруG.", "tAры", 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Count(tc.input, tc.inputsub); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
