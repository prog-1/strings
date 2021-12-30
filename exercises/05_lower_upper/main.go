package main

import (
	"bufio"
	"fmt"
	"os"
)

func ToLower(s string) string {
	var r []rune
	for _, c := range s {
		if 'A' <= c && c <= 'Z' {
			r = append(r, 'a'+c-'A')
		} else {
			r = append(r, c)
		}
	}
	return string(r)
}

func ToUpper(s string) string {
	var r []rune
	for _, c := range s {
		if 'a' <= c && c <= 'z' {
			r = append(r, 'A'+c-'a')
		} else {
			r = append(r, c)
		}
	}
	return string(r)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Print("All lower cases: ")
	fmt.Printf("%q -> %q\n", s, ToLower(s))
	fmt.Print("All upper cases: ")
	fmt.Printf("%q -> %q\n", s, ToUpper(s))
}
