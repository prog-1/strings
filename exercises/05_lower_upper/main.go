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
		for _, i := range s {
			r = append(r, unicode.ToLower(i))
			i++
		}
	}
	return string(r)
}

func ToUpper(s string) string {
	var r []rune
	if len(s) >= 1 {
		for _, i := range s {
			r = append(r, unicode.ToUpper(i))
			i++
		}
	}
	return string(r)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("All lower cases:", ToLower(s))
	fmt.Println("All upper cases:", ToUpper(s))
}
