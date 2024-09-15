package main

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()

	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()

	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)

	return 1
}

func main() {
	fmt.Println(f1()) // 1 | f1 通过 defer 修改了命名返回值 r，因此其更改影响了函数的返回值。
	fmt.Println(f2()) // 5 | f2 中 defer 修改了局部变量 t，但这没有影响返回值，因为 t 和返回值 r 是分开的。
	fmt.Println(f3()) // 1 | f3 通过值传递将 r 传入 defer，因此 defer 中对 r 的修改不影响函数的最终返回值。
	//	在函数 f3 中，通过 defer 语句传入了一个值，这种情况下，修改只会影响传入的那个值的副本，并不会影响函数外部的原始变量。
}
