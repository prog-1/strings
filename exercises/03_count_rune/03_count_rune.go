package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountRune(s string, r rune) int {
	var count int
	for i := range []rune(s) {
		if []rune(s)[i] == r {
			count++
		}
	}
	return count
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text:")
	s, _ := reader.ReadString('\n')
	var r rune
	fmt.Println("Enter rune:")
	r, _, _ = reader.ReadRune() // idea of ​​a line from someone else's code, because can't understand how to write
	fmt.Println(CountRune(s, r))
}
