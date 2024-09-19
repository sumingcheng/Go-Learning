package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("当前工作目录:", cwd)

	// 打开名为 "file" 的文件
	f, err := os.Open("file.txt")
	// 如果打开文件时发生错误，打印错误并退出
	if err != nil {
		fmt.Println(err)
		return
	}

	// 确保在函数结束时关闭文件
	defer f.Close()

	// 读取文件内容
	b, err := io.ReadAll(f)
	// 如果读取文件时发生错误，打印错误并退出
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将读取的字节转换为字符串并打印
	fmt.Println(string(b))
}
