package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountRune(s string, r rune) (cnt int) {
	for _, v := range s {
		if v == r {
			cnt++
		}
	}
	return cnt
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Enter a rune: ")
	var r rune
	fmt.Scan(&r)
	fmt.Println(CountRune(s, r))
}
