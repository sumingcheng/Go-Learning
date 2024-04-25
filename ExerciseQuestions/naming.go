package main

func main() {

}

func naming(x, y int) (sum int, err error) {
	return x + y, nil
}

// NamingSum 返回两个整数的和
func namingSum(x, y int) (sum int, err error) {
	sum = x + y // 直接使用命名返回值
	// 可以进行错误处理或其他逻辑
	return // 自动返回 sum 和 err 的当前值
}
