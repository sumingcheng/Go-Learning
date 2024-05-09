package main

import "fmt"

type N1 int

func (n N1) test() {
	fmt.Println(n)
}

func main() {
	var n N1 = 10

	fmt.Println(n) // 10

	n++
	f1 := N1.test
	f1(n) // 11

	n++
	f2 := (*N1).test
	f2(&n) // 12
}
