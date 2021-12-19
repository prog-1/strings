package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountByte(s string, b byte) (cnt int) {
	for _, r := range s {
		r_b := byte(r)
		if r_b == b {
			cnt++
		}
	}
	return cnt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Enter a byte: ")
	var b byte
	fmt.Scan(&b)
	fmt.Println(CountByte(s, b))
}
