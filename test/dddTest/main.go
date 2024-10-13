package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return `Hello, my name is ` + p.Name + ` and I am ` + strconv.Itoa(p.Age) + ` years old.`
}

func main() {
	p := Person{"Tom", 20}
	str := p.Greet()
	fmt.Println(str)
}
