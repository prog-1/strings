package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func ToLower(l string) string { // https://pkg.go.dev/unicode#ToLower // I used this unicode and found it from go.dev
	var r []rune
	for _, s := range l {
		r = append(r, unicode.ToLower(s))
	}
	return string(r)
}

func ToUpper(u string) string { // https://pkg.go.dev/unicode#ToUpper // I used this unicode and found it from go.dev
	var r []rune
	for _, s := range u {
		r = append(r, unicode.ToUpper(s))
	}
	return string(r)
}

func main() {
	fmt.Println(`Program ensures a string contains English alphabet letters
	all lower, and all upper cases`)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Result:")
	fmt.Println(ToLower(s), "lower")
	fmt.Println(ToUpper(s), "upper")
}
