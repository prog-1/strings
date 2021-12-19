package main

import (
	"bufio"
	"fmt"
	"os"
)

func ToUpper(s string) string {
	var w string
	for _, letter := range s {
		if byte(letter) >= 97 && byte(letter) <= 122 {
			letter := []byte{byte(letter) - 32}
			w += string(letter)
		} else {
			w += string(letter)
		}

	}
	return w
}
func ToLower(s string) string {
	var w string
	for _, letter := range s {
		if byte(letter) >= 65 && byte(letter) <= 90 {
			letter := []byte{byte(letter) + 32}
			w += string(letter)
		} else {
			w += string(letter)
		}

	}
	return w
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	s, _ := reader.ReadString('\n')
	fmt.Println(ToLower(s))
	fmt.Println(ToUpper(s))
}
