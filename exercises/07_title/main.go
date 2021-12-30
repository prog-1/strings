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

func Title(s string) string {
	var r []rune
	wasSep := true
	for _, c := range s {
		if IsSep(c) {
			wasSep = true
		} else if wasSep {
			wasSep = false
			r = append(r, unicode.ToTitle(c))
			continue
		}
		r = append(r, unicode.ToLower(c))

	}
	return string(r)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Printf("%q\n", Title(s))
}
