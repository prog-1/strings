package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func IsSep(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) }
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
	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(r)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(CountWords(s))
}
