package main

import (
	"context"
	proto2 "goLearning/protobuf/basic2/proto"
	"google.golang.org/grpc"
	"net"
)

// MyInfo 结构体
type MyInfo struct {
	proto2.UnimplementedMyInfoServiceServer
}

// GetData 获取数据
func (m *MyInfo) GetData(ctx context.Context, req *proto2.MyInfoRequest) (*proto2.MyInfoResponse, error) {
	return &proto2.MyInfoResponse{
		Name:       req.Name,
		Age:        req.Age,
		IsMarriage: true,
	}, nil
}

func main() {
	// 创建一个 gRPC 服务对象
	gServer := grpc.NewServer()
	// 注册服务
	proto2.RegisterMyInfoServiceServer(gServer, &MyInfo{})
	// 监听端口
	listen, _ := net.Listen("tcp", ":8080")
	// 启动gPCG服务
	_ = gServer.Serve(listen)
}
