package main

import (
	"fmt"
)

func main() {
	a, _, c := abc()
	fmt.Println(a, c)
}

func abc() (int, int, string) {
	return 1, 2, "abc"
}
