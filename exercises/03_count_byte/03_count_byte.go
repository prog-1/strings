package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountByte(s string, b byte) int {
	var count int
	for i := range []byte(s) {
		if []byte(s)[i] == b {
			count++
		}
	}
	return count
}
func main() { // code from Telegram-chat, because same problem in terminal
	fmt.Print("Enter a string: ")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(scanner.Scan(), scanner.Err(), scanner.Text())
	s := scanner.Text()
	fmt.Println("Enter a byte: ")
	fmt.Println(scanner.Scan(), scanner.Err(), scanner.Text())
	bs := scanner.Text()
	var b byte
	fmt.Sscan(bs, &b)
	fmt.Println(b)
	fmt.Println(CountByte(s, b))
}
