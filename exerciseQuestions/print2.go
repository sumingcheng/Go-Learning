package main

import "fmt"

func main() {
	i := -5
	j := 0
	k := 1

	// %+d：这个格式指示符告诉 fmt.Printf 函数
	// 不仅要输出数字，而且还要为正数前加上 + 号，负数前自然会有 - 号
	fmt.Printf("%+d %+d %+d", i, j, k) // 输出：-5 +0
}
