package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	trimSet = " \r\n\t"
)

func InSet(b byte, set string) bool {
	for _, v := range set {
		if v == rune(b) {
			return true
		}
	}
	return false
}

func TrimRight(str string) string {
	for i := len(str) - 1; i > 0 && InSet(str[i-1], trimSet); i-- {
		str = str[:i-1]
	}
	return str
}

func main() {

	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	fmt.Printf("%q\n", TrimRight(text))
}
