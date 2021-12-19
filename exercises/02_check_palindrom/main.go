package main

import "fmt"

func IsPalindrom(s string) bool {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Print("Enter text: ")
	fmt.Scan(&s)
	if IsPalindrom(s) {
		fmt.Println("String is palindrom")
	} else {
		fmt.Println("String is not palindrom")
	}
}
