package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// 线程池
var threadPool = sync.Pool{
	New: func() interface{} {
		return &thread{}
	},
}

type thread struct{}

func (t *thread) handle(conn net.Conn) {
	// 处理客户端请求
	fmt.Println("Handling request from", conn.RemoteAddr())
	// 假设处理请求需要一些时间
	time.Sleep(100 * time.Millisecond)
	conn.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		t := threadPool.Get().(*thread) // 从线程池获取一个线程
		go t.handle(conn)               // 使用线程处理请求
		threadPool.Put(t)               // 将线程放回线程池
	}
}
