package main

import "fmt"

func main() {
	//var a interface{} = 1
	//a = a.(int) + 1
	//fmt.Println(a)

	res := Plus(1, 2)
	fmt.Println(res)
}

// Go 1.17
func Plus(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		_a, _ := a.(int)
		_b, _ := b.(int)
		return _a + _b
	case string:
		_a, _ := a.(string)
		_b, _ := b.(string)
		return _a + _b
	case float32:
		_a, _ := a.(float32)
		_b, _ := b.(float32)
		return _a + _b
	default:
		panic("类型不支持")
	}
}
