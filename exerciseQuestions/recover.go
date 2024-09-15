package main

import (
	"fmt"
)

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r)
		}
	}()

	mayPanic()
	fmt.Println("After mayPanic()")
}
