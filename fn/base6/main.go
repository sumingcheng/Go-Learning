package main

import "fmt"

func main() {
	map1 := map[string]interface{}{"a": 1, "b": 2}
	map2 := map[string]interface{}{"1": "a", "2": "b"}

	// 创建一个切片，元素类型为 map[string]interface{}, 初始长度为 2
	map3 := make([]map[string]interface{}, 2)

	map3 = append(map3, map1, map2)

	forEachPrintSlice(map3)
}

func forEachPrintSlice[E any](s []E) {
	for _, e := range s {
		fmt.Println(e)
	}

	i := len(s)
	fmt.Println("len:", i)
}
