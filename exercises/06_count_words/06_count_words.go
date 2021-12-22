package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func IsSep(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) }

func countWords(s string) (cnt int) {
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			wasSep = true
		} else if wasSep {
			wasSep = false
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println("Program counts words in a string")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Printf("%v\n", countWords(text))

}
