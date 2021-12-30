package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountRune(s string, r rune) int {
	var count int
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	fmt.Print("Enter string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Print("Enter rune: ")
	var r rune
	fmt.Fscanf(reader, "%c", &r)
	if count := CountRune(s, r); count > 0 {
		fmt.Printf("%q contains rune %q (code %d) %d times.\n", s, r, r, count)
	} else {
		fmt.Printf("%q does not contain rune %q (code %d).\n", s, r, r)
	}
	// Another version:
	//
	// if count := CountRune(s, r); count > 0 {
	//	   fmt.Printf("%[1]q contains rune %[2]q (code %[2]d) %[3]d times.\n", s, r, count)
	// } else {
	// 	   fmt.Printf("%[1]q does not contain rune %[2]q (code %[2]d).\n", s, r)
	// }
}
