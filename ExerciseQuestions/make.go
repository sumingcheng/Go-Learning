package main

import "fmt"

func main() {
	// 使用 make() 函数创建一个空 map 格式如下：
	// make(map[key-type]val-type)
	s := make(map[string]int)

	delete(s, "h")

	fmt.Println(s["h"])
}
