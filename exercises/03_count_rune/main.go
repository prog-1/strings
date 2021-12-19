package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountRune(s string, b rune) int {
	var a int
	c := []rune(s)
	for i := range c {
		if c[i] == b {
			a++
		}
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text:")
	s, _ := reader.ReadString('\n')
	var b rune
	fmt.Println("Enter rune:")
	b, _, _ = reader.ReadRune()
	fmt.Println(CountRune(s, b), "times")
}
