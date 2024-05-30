package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	fmt.Println("cap", cap(s))

	a := s[:0]
	b := s[:2]

	// cap = 3 - 1 = 2
	c := s[1:2:cap(s)]

	fmt.Println(a, b, c)
}
