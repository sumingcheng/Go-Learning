package main

import "fmt"

func f4(n int) (r int) {
	defer func() {
		r += n

		recover()
	}()

	var f4 func()

	defer f4()

	f4 = func() {
		r += 2
	}

	f4()

	return n + 1
}

func main() {
	fmt.Println(f4(3))
}
