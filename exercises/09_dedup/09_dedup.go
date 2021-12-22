package main

import (
	"bufio"
	"fmt"
	"os"
)

func Dedup(s string) string {
	var r []rune
	var a rune
	for _, w := range s {
		if rune(w) != a {
			r = append(r, w)
		}
		a = rune(w)
	}
	return string(r)
}

func main() {
	fmt.Println(`Program removes duplicates by dropping
	consequent characters if they are`)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the string:")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(Dedup(s))
}
