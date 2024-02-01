package main

import (
	"context"
	"goLearning/protobuf/proto"
	"google.golang.org/grpc"
	"net"
)

// MyInfo 结构体
type MyInfo struct {
	proto.UnimplementedMyInfoServiceServer
}

// GetData 获取数据
func (m *MyInfo) GetData(ctx context.Context, req *proto.MyInfoRequest) (*proto.MyInfoResponse, error) {
	return &proto.MyInfoResponse{
		Name:       req.Name,
		Age:        req.Age,
		IsMarriage: true,
	}, nil
}

func main() {
	// 创建一个 gRPC 服务对象
	gServer := grpc.NewServer()
	// 注册服务
	proto.RegisterMyInfoServiceServer(gServer, &MyInfo{})
	// 监听端口
	listen, _ := net.Listen("tcp", ":8080")
	// 启动gPCG服务
	_ = gServer.Serve(listen)
}
