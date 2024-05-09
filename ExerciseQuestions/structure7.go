package main

type T struct {
	n int
}

func (t *T) Set(n int) {
	t.n = n
}

func getT() *T {
	return &T{}
}

func main() {
	// 要调用定义在T类型上的Set方法，需要使用指针*T。
	// 这是因为Set方法的接收者是指针类型*T，而不是值类型T。
	getT().Set(1)
}
