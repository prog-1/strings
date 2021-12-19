package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func ToLower(s string) string {
	var tmp []rune
	for _, r := range s {
		tmp = append(tmp, unicode.ToLower(r))
	}
	return string(tmp)
}

func ToUpper(s string) string {
	var tmp []rune
	for _, r := range s {
		tmp = append(tmp, unicode.ToUpper(r))
	}
	return string(tmp)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(s)
	fmt.Println(ToLower(s), ToUpper(s))
}
