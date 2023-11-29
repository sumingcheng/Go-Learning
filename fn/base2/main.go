package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile()
	fmt.Println("文件读取完毕")
}

func readFile() {
	// 捕获错误
	defer catchPanic()
	file, err := os.Open("./fn/base2/1.tx1t")

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("关闭文件失败")
		}
	}(file)

	if err != nil {
		panic("打开文件失败")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		println(scanner.Text())
	}
}

func catchPanic() {
	if err := recover(); err != nil {
		fmt.Println("捕获到异常", err)
	}
}
