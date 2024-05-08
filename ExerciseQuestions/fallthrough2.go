package main

import "fmt"

func main() {
	match := func(i int) bool {
		switch i {
		case 1, 2:
			return true
		}
		return false
	}

	fmt.Println(match(1)) // 输出 true
	fmt.Println(match(2)) // 输出 true
}
