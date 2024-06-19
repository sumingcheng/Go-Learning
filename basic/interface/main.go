package main

import "fmt"

// Display 接口，要求实现 Show 方法。
type Display interface {
	Show(message string) error
}

// TextPrinter 结构体，包含要打印的 Content。
type TextPrinter struct {
	Content string
}

// Show 方法，实现 Display 接口。打印 Content，返回 nil。
func (tp TextPrinter) Show(message string) error {
	fmt.Println(tp.Content) // 实际打印内容为 Content 字段。
	return nil              // 假设总是成功，返回 nil。
}

func main() {
	text := TextPrinter{"Hello, Interface!"} // 创建 TextPrinter 实例。
	var display Display = text               // 通过 Display 接口调用 Show。
	err := display.Show("123")

	if err != nil {
		return // 如果有错误，提前返回。
	}
}
