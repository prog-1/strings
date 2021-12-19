package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func IsSep(r rune) bool { return unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) }

func Title(s string) []string {
	var w []string
	var tmp, v []rune
	wasSep := true
	for _, r := range s {
		if IsSep(r) {
			if !wasSep {
				for _, i := range tmp {
					i = (unicode.To(unicode.LowerCase, i))
					v = append(v, i)
					v[0] = unicode.ToTitle(v[0])
				}
				w = append(w, string(v))
				tmp, v = nil, nil
			}
			wasSep = true
		} else {
			wasSep = false
			tmp = append(tmp, r)
		}
	}
	if len(tmp) > 0 {
		for _, i := range tmp {
			i = (unicode.To(unicode.LowerCase, i))
			v = append(v, i)
			v[0] = unicode.ToTitle(v[0])
		}
		w = append(w, string(v))
	}
	return w
}
func main() {
	r := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(r)
	fmt.Print("Enter a string: ")
	scanner.Scan()
	s := scanner.Text()

	fmt.Println(Title(s))
}
