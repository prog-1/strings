package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func ToLower(s string) string {
	var r []rune
	if len(s) >= 1 {
		for _, z := range s {
			r = append(r, unicode.ToLower(z))
			z++
		}
	}
	return string(r)
}

func ToUpper(s string) string {
	var r []rune
	if len(s) >= 1 {
		for _, z := range s {
			r = append(r, unicode.ToUpper(z))
			z++
		}
	}
	return string(r)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(ToLower(s), ToUpper(s))
}
