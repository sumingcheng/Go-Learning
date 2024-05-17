package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	// r数组的值与最终a被修改后的值一致了。
	// 这个例子中我们使用了*[5]int作为range表达式，其副本依旧是一个指向原数组 a的指针，
	// 因此后续所有循环中均是&a指向的原数组亲自参与的，因此v能从&a指向的原数组中取出a修改后的值。
	for i, v := range &a {
		if i == 0 {
			a[1] = 22
			a[2] = 33
		}

		r[i] = v
	}

	fmt.Println("a = ", a)
	fmt.Println("r = ", r)
}
