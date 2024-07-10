package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"John Doe", 30}
	fmt.Printf("Using %%v: %v\n", p)
	fmt.Printf("Using %%+v: %+v\n", p)
	fmt.Printf("Using %%#v: %#v\n", p)
}
