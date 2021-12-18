package main

import (
	"bufio"
	"fmt"
	"os"
)

func IsPalindrom(s string) bool {
	r := []rune(s)
	i, j := 0, len(r)-1
	for ; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("%v\n", IsPalindrom(text))
}
