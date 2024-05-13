package main

import "fmt"

func main() {
outerLoop:
	for i := 1; i < 3; i++ {
		for j := 2; j < 3; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				break outerLoop
			}
		}
	}
	fmt.Println("Exited loop")
}
