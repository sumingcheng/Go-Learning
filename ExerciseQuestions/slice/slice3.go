package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}
	newSlice := original[1:3] // newSlice 的内容是 [2, 3]

	// 当前 newSlice 的容量是 4，因为它从 original 的第2个元素开始，共享到 original 的末尾。
	fmt.Println("追加前的新切片:", newSlice, "cap:", cap(newSlice))

	// append 操作导致 newSlice 的容量不足以存放更多的元素，因此会重新分配数组。
	newSlice = append(newSlice, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110)

	// 检查 original 的内容，看是否被改变。
	fmt.Println("附加到 newSlice 后的原始:", original)              // [1 2 3 4 5]
	fmt.Println("追加后的新切片:", newSlice, "cap:", cap(newSlice)) // [2 3 10 20 30 40 50 60 70 80 90 100 110]
}
