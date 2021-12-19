package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountByte(str string, sbyte byte) int {
	var count int
	for i := 0; i < len(str); i++ {
		if str[i] == sbyte {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	var sbyte byte
	fmt.Print("Enter byte: ")
	fmt.Scan(&sbyte)
	fmt.Println("Byte", sbyte, "contains", CountByte(text, sbyte), "2 times")
}
