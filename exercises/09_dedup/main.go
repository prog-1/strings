package main

import (
	"bufio"
	"fmt"
	"os"
)

func Dedup(s string) string {
	var r []rune
	var prev rune
	for _, c := range s {
		if c == '\x00' {
			r = append(r, c)
		}

		if c != prev {
			r = append(r, c)
		}
		prev = c
	}
	return string(r)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter text: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Printf("%q\n", Dedup(s))
}
