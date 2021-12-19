package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func Title(s string) string {
	var outputr []rune
	tmp := 1
	for _, r := range s {
		if tmp == 1 {
			outputr = append(outputr, unicode.ToUpper(r))
		} else if tmp > 1 {
			outputr = append(outputr, unicode.ToLower(r))
		}
		tmp++
	}
	return string(outputr)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(Title(s))
}
