package main

import "fmt"

func main() {
	originalSlice := []int{1, 2, 3}
	modifySlice(originalSlice)
	fmt.Println(originalSlice) // 输出 [100, 2, 3]
	fmt.Printf("%p\n", originalSlice)
}

func modifySlice(s []int) {
	s[0] = 100
	fmt.Printf("%p\n", s)
}
