package main

import "fmt"

func main() {
	//type Test struct {
	//	a int8  // 1
	//	c bool  // 1
	//	b int32 // 4
	//}
	//
	//test := Test{1, true, 2}
	//
	//fmt.Println(unsafe.Sizeof(test)) // 8

	//type Logistics struct {
	//	orderId     string
	//	productName string
	//	address     struct {
	//		province string
	//		city     string
	//		district string
	//		detail   string
	//	}
	//}

	type Address struct {
		province string
		city     string
		district string
		detail   string
	}

	type LogisticsInfo struct {
		orderId     string
		productName string
		address     Address
	}

	//myAddress := Address{
	//	province: "广东省",
	//	city:     "深圳市",
	//	district: "南山区",
	//	detail:   "科技园",
	//}
	//
	//logisticsInfo := LogisticsInfo{
	//	orderId:     "123456789",
	//	productName: "iPhone 12",
	//	address:     myAddress,
	//}
	//
	//fmt.Printf("%+v\n", logisticsInfo)

	// 匿名结构体

	//logisticsInfo := struct {
	//	orderId     string
	//	productName string
	//}{
	//	orderId:     "123456789",
	//	productName: "iPhone 12",
	//}
	//
	//fmt.Println(logisticsInfo)

	logisticsInfo := LogisticsInfo{
		"123456789",
		"iPhone 12",
		Address{
			province: "广东省",
			city:     "深圳市",
			district: "南山区",
			detail:   "科技园",
		},
	}

	fmt.Print(logisticsInfo)
}
