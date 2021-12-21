package main

import (
	"bufio"
	"fmt"
	"os"
)

func Count(s string) int {}

func main() {
	fmt.Println(`Program finds how many times a substring exists in a string with
	overlaps`)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the string")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Enter the substring")
	scanner.Scan()
	ss := scanner.Text()
	fmt.Println(Count(s))
}
