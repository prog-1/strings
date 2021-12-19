package main

import (
	"bufio"
	"fmt"
	"os"
)

func testEq(a, b []slice) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(r)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Print("Enter a substring: ")
	scanner.Scan()
	sub := scanner.Text()
	if i := substr(sub, s); i >= 0 {
		fmt.Printf("The substring %q occurs in  the string %q at index %d.", sub, s, i)
	} else {
		fmt.Printf("The substring %q does not occur in  the string %q", sub, s)
	}
}
func substr(sub, s string) int {
	for l := range s {

	}

	return -1
}
