package main

import "fmt"

func InSet(b byte, set string) bool {
	for _, v := range set {
		if v == rune(b) {
			return true
		}
	}
	return false
}

const trimSet = " \r\n\t"

func TrimLeft(s string) string {
	for len(s) > 0 && InSet(s[0], trimSet) {
		s = s[1:]
	}
	return s
}

func main() {
	fmt.Printf("%q\n", TrimLeft("   \r\r\r\n\t    hello"))
}
