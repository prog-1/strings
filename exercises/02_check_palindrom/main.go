package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsPalindrome(s string) bool {
	var reversed string

	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])

	}
	for i := range s {
		if s[i] != reversed[i] {
			return false
		}
	}
	return true
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Program  checks whether a given string is a palindrom")
	fmt.Println("Enter the string")
	scanner.Scan()
	s := scanner.Text()
	if !IsPalindrome(s) {
		fmt.Print("String is not a palindrome")

	} else {
		fmt.Print("String is a palindrome")
	}
}
