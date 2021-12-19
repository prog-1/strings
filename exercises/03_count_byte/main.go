package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountByte(s string, b byte) int {
	var a int
	c := []byte(s)
	fmt.Println(c)
	for i := range c {
		if c[i] == b {
			a++
		}
	}
	return a
}
func main() {
	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(scanner.Scan(), scanner.Err(), scanner.Text())
	s := scanner.Text()
	fmt.Println("Enter a byte: ")
	fmt.Println(scanner.Scan(), scanner.Err(), scanner.Text())
	bs := scanner.Text()
	var b byte
	fmt.Sscan(bs, &b)
	fmt.Println(b)
	fmt.Println(CountByte(s, b), "times")
}
