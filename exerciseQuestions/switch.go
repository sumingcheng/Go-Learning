package main

func alwaysFalse() bool {
	return false
}

func main() {
	// 注意
	switch alwaysFalse(); {
	case true:
		println(true) // 如果返回 true，打印 "true"
	case false:
		println(false) // 如果返回 false，打印 "false"
	}
}
