package main

import "fmt"

func modifySlice(slice []int) {
	slice[0] = 100
}

func main() {
	mySlice := []int{1, 2, 3}
	fmt.Println("Before:", mySlice) // 输出初始值

	modifySlice(mySlice)           // 传递切片
	fmt.Println("After:", mySlice) // 输出修改后的值
}
