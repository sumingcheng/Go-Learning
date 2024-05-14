package main

import "fmt"

type foo3 struct{ Val int }
type bar3 struct{ Val int }

func main() {
	a := &foo3{Val: 5}
	b := &foo3{Val: 5}
	c := foo3{Val: 5}

	d := bar3{Val: 5}
	e := bar3{Val: 5}
	f := bar3{Val: 5}

	fmt.Print(a == b, c == foo3(d), e == f) // 输出比较结果：
	// a == b 比较两个指针（不同实例，结果为 false）
	// c == foo3(d) 将 d (bar3 类型) 转换为 foo3 类型并比较（编译错误，类型不匹配）
	// e == f 比较两个 bar3 实例（相同内容，结果为 true）
}
