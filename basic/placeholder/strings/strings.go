package strings

import (
	"fmt"
	"strings"
)

func Str() {
	var builder strings.Builder

	// 追加字符串
	builder.WriteString("Hello, ")
	builder.WriteString("World!")
	builder.WriteRune('\n')

	// 获取最终的字符串
	result := builder.String()

	// 输出
	fmt.Println(result)

	// 重置 builder
	builder.Reset()

	// 再次使用
	builder.WriteString("Goodbye, World!")
	result = builder.String()
	fmt.Println(result)

}
