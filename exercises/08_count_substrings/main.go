package main

import (
	"bufio"
	"fmt"
	"os"
)

func Count(s, sub string) int {
	var count int
	x := len(sub)
	for c := range s {
		if x+c > len(s) {
			break
		}
		if sub == s[c:x+c] {
			count++
		}
	}
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Print("Enter substring: ")
	scanner.Scan()
	sub := scanner.Text()
	if count := Count(s, sub); count > 0 {
		fmt.Printf("%q contains substring %q %d times.\n", s, sub, count)
	} else {
		fmt.Printf("%q does not contain substring %q.\n", s, sub)
	}
}
