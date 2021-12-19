package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountRune(str string, sbyte rune) int {
	var count int
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		if str1[i] == sbyte {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	var srune rune
	fmt.Print("Enter rune: ")
	srune, _, _ = reader.ReadRune()
	fmt.Println("Rune", srune, "contains", CountRune(text, srune), "time(s)")
}
