package main

import "fmt"

func add(args ...int) int {
	sum := 0

	for _, arg := range args {
		sum += arg
	}

	return sum
}

func main() {
	s1 := add(1, 2, 3)
	fmt.Println(s1)

	s2 := add([]int{3, 3, 3}...)
	fmt.Println(s2)
}
