package main

import (
	"bufio"
	"fmt"
	"os"
)

func Dedup(s string) string {
	var z []rune
	var zz rune
	for _, r := range s {
		if rune(r) != zz {
			z = append(z, r)
		}
		zz = rune(r)
	}
	return string(z)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(Dedup(s))
}
