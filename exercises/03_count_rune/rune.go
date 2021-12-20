package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountRune(s string, z rune) int {
	var i int
	j := []rune(s)
	for a := range j {
		if j[a] == z {
			i++
		}
	}
	return i
}

func main() {
	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(r)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()
	fmt.Print("Enter a rune: ")
	var z rune
	fmt.Fscanf(r, "%c", &z)
	fmt.Printf("s = %[1]q\nb = %[2]c (%[2]d)", s, z)
	fmt.Println(s, "contains", z, CountRune(s, z), "times")
}
