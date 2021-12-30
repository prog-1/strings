package main

import (
	"testing"
)

func TestDedup(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"leading ascii 0", "\x00", "\x00"},
		{"just one letter", "a", "a"},
		{"string without the same consequent separators", "*-+!?,.:;'/<>(){}[]%", "*-+!?,.:;'/<>(){}[]%"},
		{"string with the same consequent separators", "!!!!???++++!!", "!?+!"},
		{"number without the same consequent characters", "1234567891", "1234567891"},
		{"number with the same consequent characters", "00011111000111", "0101"},
		{"words without the same consequent characters", "how are you", "how are you"},
		{"words with the same consequent characters", "good morning and hello", "god morning and helo"},
		{"words with the same consequent characters + separators", "Merry christmas!!!", "Mery christmas!"},
		{"unicode without the same consequent characters", "привет, мир!", "привет, мир!"},
		{"unicode with the same consequent characters", "доброе утрооо! чтооо делаешь???", "доброе утро! что делаешь?"},
		{"words in spaces with the same consequent characters", "  he's too funny  today   ", " he's to funy today "},
		{"mixed unicode with the same consequent characters", "čeлovekk ии ssoбаka", "čeлovek и soбаka"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Dedup(tc.input); got != tc.want {
				t.Errorf("got = %q, want = %q", got, tc.want)
			}
		})
	}
}
