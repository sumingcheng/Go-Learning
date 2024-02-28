package main

import (
	"context"
	"fmt"
	proto2 "goLearning/protobuf/basic2/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 拨号连接
	conn, _ := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// 实例化 gRPC 客户端
	client := proto2.NewMyInfoServiceClient(conn)
	// 调用方法
	data, _ := client.GetData(context.Background(), &proto2.MyInfoRequest{
		Name:       "张三",
		Age:        18,
		IsMarriage: false,
	})
	// 取出数据
	fmt.Println(data)
}
