package main

import (
	"bufio"
	"fmt"
	"os"
)

func Count(s, sub string) int {
	var cnt int

	return cnt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()

	fmt.Println("Enter a substring: ")
	scanner.Scan()
	sub := scanner.Text()

	if i := Count(s, sub); i > 0 {
		fmt.Printf("String %q contains %q %v times", s, sub, i)
	} else {
		fmt.Printf("String %q does not contains %q", s, sub)
	}
}
