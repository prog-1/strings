package main

import (
	"bufio"
	"fmt"
	"os"
)

func Dedup(s string) string {
	var tmp []rune
	var a rune
	for _, r := range s {
		if rune(r) == a {

		} else {
			tmp = append(tmp, r)
			a = rune(r)
		}

	}
	return string(tmp)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Program removes duplicates")
	fmt.Println("Enter the string")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(Dedup(s))
}
