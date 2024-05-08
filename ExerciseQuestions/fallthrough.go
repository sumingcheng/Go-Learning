package main

import "fmt"

func main() {
	isMatch := func(i int) bool {
		switch i {
		case 1:
			//fallthrough
		case 2:
			return true
		}
		return false
	}

	fmt.Println(isMatch(1)) // 输出 true
	fmt.Println(isMatch(2)) // 输出 true
}
