package main

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

	//a := true
	//aType := reflect.TypeOf(a)
	//fmt.Println(aType) // bool
	//
	//b := 1
	//bType := reflect.ValueOf(b).Kind()
	//fmt.Println(bType) // int

	//	——————————————————

	/*	var s = "123"
		a, err := strconv.Atoi(s)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(a) // 字符串转换成 int 类型*/

	//	——————————————————
	/*	var a = 123
		str := strconv.Itoa(a)

		fmt.Println(str)*/

	//	——————————————————
	//字符串转float

	//f, err := strconv.ParseFloat("3.1415926", 64)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(f)

	//	——————————————————

	/*	i, err := strconv.ParseInt("123", 10, 64)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(i)*/

	//	——————————————————

	/*s := strconv.FormatFloat(123.45, 'f', 1, 64)
	fmt.Println(s)*/

}
