package main

import "fmt"

const (
	a11 = iota
	b2  = iota
)

const (
	name = "name"
	c3   = iota
	d4   = iota
)

func main() {
	fmt.Println(a11)
	fmt.Println(b2)
	fmt.Println(c3)
	fmt.Println(d4)
}
