package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{1, 2, 3, 4, 5}

	if !sort.IntsAreSorted(data) {
		sort.Ints(data)
		fmt.Println("数据已排序。")
	} else {
		fmt.Println("数据已经是排序状态。")
	}
}
