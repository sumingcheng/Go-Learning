package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var a int8 = 123
	//var b = uint8(a)
	//fmt.Println(b)

	//var a int8 = -123
	//var b = uint8(a)
	//// 256 - 123 = 133
	//fmt.Println(b) // 133

	//	——————————————————

	//var a int32 = 123
	//var b int64 = int64(a)
	//fmt.Println(b)

	//	——————————————————

	a := true
	aType := reflect.TypeOf(a)
	fmt.Println(aType) // bool

	b := 1
	bType := reflect.ValueOf(b).Kind()
	fmt.Println(bType) // int
}
