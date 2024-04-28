package main

import "fmt"

func main() {
	v := []int{1, 2, 3, 4}

	// index item
	for _, item := range v {
		v = append(v, item)
	}

	fmt.Println(v)
}
