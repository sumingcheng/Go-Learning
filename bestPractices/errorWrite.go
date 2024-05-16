package main

import (
	"fmt"
	"io"
)

// 定义一个结构体，包括一个io.Writer接口和一个错误对象，用于在写入过程中捕获错误
type errWriter struct {
	io.Writer
	err error
}

// Write方法实现了io.Writer接口。这个方法在err不为nil时停止写入，保证不会在出现错误后继续操作。
func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		// 如果已存在错误，停止执行并返回错误
		return 0, e.err
	}
	var n int
	// 尝试写入数据，如果写入过程中出现错误，将错误保存在e.err中
	n, e.err = e.Writer.Write(buf)
	return n, nil // 返回写入的字节数和nil（错误已经被捕获）
}

func WriteResponse(
	w io.Writer,
	st Status,
	headers []Header,
	body io.Reader,
) error {
	ew := &errWriter{Writer: w}
	// 写入HTTP状态行
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)

	// 遍历并写入所有HTTP头部
	for _, h := range headers {
		fmt.Fprintf(ew, "%s: %s\r\n", h.Key, h.Value)
	}

	// 写入头部与主体的分隔符
	fmt.Fprintf(ew, "\r\n")
	// 从body中复制数据到Writer，使用ew进行错误管理
	io.Copy(ew, body)

	// 返回在写入过程中捕获的任何错误
	return ew.err
}
