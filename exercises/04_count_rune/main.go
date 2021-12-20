package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountRune(s string, r rune) int {
	var i int
	j := []rune(s)
	for a := range j {
		if j[a] == r {
			i++
		}
	}
	return i
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	fmt.Println("Enter string:")
	scanner.Scan()
	s := scanner.Text()
	fmt.Println("Enter rune:")
	var r rune
	fmt.Fscanf(reader, "%c", &r)
	if cnt := CountRune(s, r); cnt > 0 {
		fmt.Printf("%[1]q contains rune %[2]q (code %[2]d) %[3]d times.\n", s, r, cnt)
	} else {
		fmt.Printf("%[1]q does not contain rune %[2]q (code %[2]d).\n", s, r)
	}
}
