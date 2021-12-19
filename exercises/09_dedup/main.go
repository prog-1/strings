package main

import (
	"bufio"
	"fmt"
	"os"
)

func Dedup(s string) string {
	var outputr []rune
	var prev rune
	for _, r := range s {
		if r != prev {
			outputr = append(outputr, r)
		}
		prev = r
	}
	return string(outputr)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(Dedup(s))
}
