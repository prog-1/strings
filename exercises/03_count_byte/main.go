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
	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text:")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Enter byte:")
	var b byte
	fmt.Fscanf(r, "%c", &b)
	if cnt := CountByte(s, b); cnt > 0 {
		fmt.Printf("%[1]q contains byte %[2]q (code %[2]d) %[3]d times.\n", s, b, cnt)
	} else {
		fmt.Printf("%[1]q does not contain byte %[2]q (code %[2]d).\n", s, b)
	}
}
