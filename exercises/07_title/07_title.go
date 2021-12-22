package main

import (
	"bufio"
	"fmt"
	"os"
)

func Title(string) string {}

func main() {
	fmt.Println(`program makes every first letter of every word capital, while
	all other letters should become lower case`)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(Title(s))
}
