package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountRune(s string, b rune) int {
	var count int
	s1 := []rune(s)
	for i := 0; i < len(s); i++ {
		if s1[i] == b {
			count++
		}

	}
	return count
}

func main() {
	fmt.Println("This program counts the number of instances of the byte in the string.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	text, _ := reader.ReadString('\n')
	var srune rune
	fmt.Println("Enter a rune: ")
	srune, _, _ = reader.ReadRune()
	fmt.Println("Rune", srune, "contains", CountRune(text, srune), "time(s)")

}
