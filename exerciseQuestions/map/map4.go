package main

import "fmt"

func main() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}

	counter := 0

	for index, item := range m {
		if counter == 0 {
			delete(m, "A")
		}

		counter++

		fmt.Println(index, item)
	}

	fmt.Println("counter is ", counter) // 2 或 3 ,因为map是无序的
}
