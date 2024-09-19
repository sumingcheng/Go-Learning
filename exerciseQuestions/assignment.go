package main

import "fmt"

func main() {
	x, y := 1, 2
	x, y = y, x
	fmt.Println(x, y)

	a, b, c, d := returnMultipleValues()
	fmt.Println(a, b, c, d)
}

func returnMultipleValues() (int, int, int, int) {
	return 1, 2, 3, 4
}
