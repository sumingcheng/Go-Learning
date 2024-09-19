package main

import "fmt"

const (
	greeting = "Hello, World!"
	first    = 1 << iota
	second
)

func main() {
	fmt.Println(first, second)
}
