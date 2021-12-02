package main

import (
	"bufio"
	"fmt"
	"os"
)

func InSet(b byte, set []byte) bool {
	for _, v := range set {
		if v == b {
			return true
		}
	}
	return false
}

var trimSet = []byte(" \r\n\t")

func TrimRight(s string) string {
	for s != "" && InSet(s[len(s)-1], trimSet) {
		s = s[:len(s)-1]
	}
	return s
}

func main() {
	fmt.Printf("%q\n", TrimRight("   \r\r\r\n\t    hello"))

	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	fmt.Printf("%q\n", TrimRight(text))
}
