package main

import "testing"

func TestCountWords(t *testing.T) {
	for _, tc := range []struct {
		name string
		str  string
		want int
	}{
		{"empty", "", 0},
		{"Hello, world", "hello, world", 2},
		{"Thats a lot of text", `Java is a language that basically outsourced memory management entirely to its garbage collector. This turned out to be a big mistake. However, to be able to explain why, I need to cover more of the details.
		Let us start at the beginning. It is 1991 and the work on Java has begun. Garbage collectors are all the rage. Research looks very promising and the designers of Java are placing their bets on advanced garbage collectors being able to solve essentially all challenges in managing memory.
		For this reason, all objects in Java are designed to be allocated on the heap, with the exception of primitive types such as integers and floating point values. When talking about memory allocation, we generally distinguish between what is called the heap and the stack. The stack is very fast to use, but has limited space and can only be used for objects that do not need to exist beyond the lifetime of a function call. It is only for local variables. The heap can be used for all objects. Java basically ignored the stack and chose to allocate everything on the heap, except primitives like integers and floats. Whenever you write new Something() in Java, you are consuming memory on the heap. However, this type of memory management is actually quite expensive in terms of memory usage. You would think that creating an object with only a 32-bit integer would only require 4 bytes of memory. `, 243}, // https://itnext.io/go-does-not-need-a-java-style-gc-ac99b8d26c60
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := CountWords(tc.str); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
