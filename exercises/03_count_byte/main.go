package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountByte(s string, b byte) int {
	var count int
	for _, c := range []byte(s) {
		if c == b {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Print("Enter byte: ")
	var b byte
	fmt.Fscanf(reader, "%c", &b)
	if count := CountByte(s, b); count > 0 {
		fmt.Printf("%q contains byte %q (code %d) %d times.\n", s, b, b, count)
	} else {
		fmt.Printf("%q does not contain byte %q (code %d).\n", s, b, b)
	}
	// Another version:
	//
	// if count := CountByte(s, b); count > 0 {
	//	   fmt.Printf("%[1]q contains byte %[2]q (code %[2]d) %[3]d times.\n", s, b, count)
	// } else {
	// 	   fmt.Printf("%[1]q does not contain byte %[2]q (code %[2]d).\n", s, b)
	// }
}
