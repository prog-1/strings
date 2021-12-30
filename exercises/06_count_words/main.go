package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func IsSep(r rune) bool {
	return unicode.IsPunct(r) || unicode.IsSpace(r) || unicode.IsSymbol(r)
}

// CountWords counts the number of words that occurs in a string `s`.
func CountWords(s string) (cnt int) {
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			wasSep = true
		} else if wasSep {
			wasSep = false
			cnt++
		}
	}
	return cnt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Printf("There are %d words in the text.\n", CountWords(s))
}
