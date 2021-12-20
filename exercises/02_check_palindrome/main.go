package main

import "fmt"

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
	var s string
	fmt.Println("Enter text:")
	fmt.Scan(&s)
	if IsPalindrom(s) {
		fmt.Printf("%q is a palindrome.\n", s)
	} else {
		fmt.Printf("%q is not a palindrome.\n", s)
	}
}
