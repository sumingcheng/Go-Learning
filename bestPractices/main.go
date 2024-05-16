package main

import (
	"os"
	"strings"
)

type Status struct {
	Code   int
	Reason string
}

type Header struct {
	Key   string
	Value string
}

func main() {
	// 创建一个示例的HTTP状态和头部
	status := Status{Code: 200, Reason: "OK"}
	headers := []Header{
		{"Content-Type", "text/plain"},
	}

	// 响应体是一个简单的字符串
	body := strings.NewReader("Hello, world!")

	// 使用os.Stdout作为Writer，这里仅为示例，通常这会是网络连接的一部分
	err := WriteResponse(os.Stdout, status, headers, body)
	if err != nil {
		// 如果有错误发生，输出错误
		println("Error writing response:", err.Error())
	}
}

// 此处省略之前代码中的errWriter和WriteResponse的定义
