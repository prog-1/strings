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
	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(r)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Print("Enter a byte: ")
	var b byte
	fmt.Fscanf(r, "%c", &b)
	fmt.Printf("s = %[1]q\nb = %[2]c (%[2]d)", s, b)
	fmt.Print(" => contains ")
	fmt.Printf("%[2]c", s, b, " ")
	fmt.Println(" ", CountByte(s, b), "times")
}
