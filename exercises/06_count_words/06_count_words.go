package main

import (
	"fmt"
	"unicode"
)

func CountWords(str string) int {
	var count int
	var wasLetter bool
	for _, v := range str {
		if (unicode.IsSpace(v) || unicode.IsPunct(v)) && wasLetter {
			wasLetter = false
		} else if !wasLetter {
			wasLetter = true
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountWords("Hello World"))
}
