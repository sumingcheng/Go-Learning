package main

import "fmt"

func main() {
	var a *int
	fmt.Println(a)
	fmt.Printf("%p\n", a)
	fmt.Println(&a)
	fmt.Println(*a)
}
