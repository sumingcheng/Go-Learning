package main

import "fmt"

var nil = new(int)

func main() {
	var p *int
	println(nil, p)
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
