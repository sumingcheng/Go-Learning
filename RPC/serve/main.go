package main

import (
	"context" // Go 的上下文库，用于获取上下文信息
	"log"     // 日志库，用于打印日志信息
	"net"     // 网络库，用于网络相关操作

	"google.golang.org/grpc" // 引入 gRPC 库
	pb "hello/proto"         // 引入 protobuf 生成的包
)

// server 用于实现 hello.proto 中定义的 Greeter 服务。
type server struct {
	pb.UnimplementedGreeterServer // 嵌入未实现的服务接口，以便我们可以填充我们需要的方法
}

// SayHello 实现了 Greeter 服务接口。
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	// 构造 HelloResponse，将接收到的 HelloRequest 中的 greeting 字段加上 "Hello " 前缀
	return &pb.HelloResponse{Reply: "Hello " + in.Greeting}, nil
}

func main() {
	// 监听本地的 50051 端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err) // 监听失败则打印错误并退出
	}

	// 创建一个 gRPC 服务器
	s := grpc.NewServer()
	// 将我们实现的 Greeter 服务注册到 gRPC 服务器中
	pb.RegisterGreeterServer(s, &server{})
	// 启动 gRPC 服务器，监听并接收客户端请求
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err) // 启动失败则打印错误并退出
	}
}
