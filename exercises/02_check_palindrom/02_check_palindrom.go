package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsPalindrom(s string) bool {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j+1 {
		if r[i] == r[j] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println("Program checks whether a given string is a palindrome")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("%v\n", IsPalindrom(text))
}
