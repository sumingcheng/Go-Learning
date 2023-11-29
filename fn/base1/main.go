package main

import "fmt"

func main() {
	//test()
	//res := test5()
	//res()

}

func test() {
	a := 10
	_test := func() {
		fmt.Println(a)
	}
	_test()
}

func test1() func() {
	return func() {
		fmt.Println("test1")
	}
}

func test2(a int, b int) func(aa int, bb int) func(aaa int, bbb int) {
	return func(aa int, bb int) func(aaa int, bbb int) {
		return func(aaa int, bbb int) {

		}
	}
}

func test3(cb1 func(a int), cb2 func(b int)) func(aa int, bb int) func(aaa int, bbb int) {
	return func(aa int, bb int) func(aaa int, bbb int) {
		return func(aaa int, bbb int) {
			cb1(1)
		}
	}
}

func test4() {
	/*
		test 函数的作用域，捆绑着 test4 函数的作用域，也捆绑着全局作用域
	*/
	//test := func() {
	//	fmt.Println("fun1")
	//}
}

func test5() func() {
	count := 0

	return func() {
		fmt.Println(count)
	}
}

func test6() {

}
