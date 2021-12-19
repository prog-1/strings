package main

import (
	"bufio"
	"fmt"
	"os"
)

func Count(s, sub string) int {
	var cnt int
	if len(s) > len(sub) {
		for i := 0; i < len(s)-len(sub); i++ {
			if s[i:i+len(sub)] == sub {
				cnt++
			}
		}
		return cnt
	}
	return 0
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Program finds how many times a substring exists in a string with overlaps.")
	fmt.Println("Enter the string")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Enter the substring")
	scanner.Scan()
	sub := scanner.Text()
	fmt.Println(Count(s, sub))
}
