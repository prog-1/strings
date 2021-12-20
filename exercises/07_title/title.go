package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

//now it works only with 1 word, but I will improve it
func Title(s string) string {
	var r []rune
	if len(s) >= 1 {
		for _, z := range s {
			r = append(r, unicode.ToLower(z))
			z++
			r[0] = unicode.ToUpper(r[0])
		}
	}
	return string(r)
}
func main() {
	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(r)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(Title(s))
}
