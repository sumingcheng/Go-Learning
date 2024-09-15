package main

import (
	"context" // Go 的上下文库，用于设置超时等上下文信息
	"log"     // 日志库，用于打印日志信息
	"time"    // 时间库，用于处理时间相关的操作

	"google.golang.org/grpc" // 引入 gRPC 库
	pb "hello/proto"         // 引入 protobuf 生成的包
)

func main() {
	// 连接到服务端，地址是本地的 50051 端口
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect.go: %v", err) // 连接失败则打印错误并退出
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn) // 函数结束时关闭连接

	// 使用连接创建 Greeter 服务的客户端
	c := pb.NewGreeterClient(conn)

	// 创建带有超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel() // 函数结束时取消上下文

	// 调用 SayHello 方法，发送 HelloRequest，并接收 HelloResponse
	r, err := c.SayHello(ctx, &pb.HelloRequest{Greeting: "world"})

	if err != nil {
		log.Fatalf("could not greet: %v", err) // 调用失败则打印错误并退出
	}
	log.Printf("Greeting: %s", r.GetReply()) // 打印收到的回复
}
