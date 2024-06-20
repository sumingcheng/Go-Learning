package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取文件示例
func main() {
	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Print(line)
				break
			} else {
				fmt.Println("读文件发生错误", err)
				return
			}
		} else {
			fmt.Print(line)
		}
	}
}
