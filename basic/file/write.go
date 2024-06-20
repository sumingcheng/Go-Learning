package main

import (
	"fmt"
	"os"
)

// os.O_RDONLY	以只读模式打开文件。
// os.O_WRONLY	以只写模式打开文件。
// os.O_RDWR	以读写模式打开文件。
// os.O_APPEND	写操作时将数据追加到文件尾部。
// os.O_CREATE	如果文件不存在，则创建该文件。
// os.O_EXCL	与 os.O_CREATE 结合使用时，文件必须不存在，否则操作将失败。
// os.O_SYNC	打开文件进行同步 I/O，确保写入的数据立即刷新到底层硬件。
// os.O_TRUNC	如果可能，打开时清空文件。
func main() {
	file, err := os.OpenFile("file.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	content := "Hello, world!\n"
	n, err := file.Write([]byte(content))

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(n, "bytes written.")
	}

}
