package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义命令行参数
	var ip = flag.String("ip", "127.0.0.1", "服务器IP地址")
	var port = flag.Int("port", 8080, "服务器端口号")
	var debug = flag.Bool("debug", false, "开启调试模式")

	// 解析命令行参数
	flag.Parse()

	// 使用解析后的参数
	fmt.Printf("服务器地址：%s\n", *ip)
	fmt.Printf("端口号：%d\n", *port)
	fmt.Printf("调试模式：%t\n", *debug)
}
