package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountByte(s string, b byte) int {
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			count++
		}

	}
	return count
}

func main() {
	fmt.Println("This program counts the number of instances of the byte in the string.")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Enter a byte: ")
	var b byte
	fmt.Scan(&b)
	fmt.Println(CountByte(s, b))
}
