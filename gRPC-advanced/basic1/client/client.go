package main

import (
	"context"
	"encoding/json"
	"fmt"
	"goLearning/gRPC-advanced/basic1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

//var md = metadata.New(map[string]string{
//	"token": "123456",
//})

var md = metadata.Pairs("token", "123456")

func main() {
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	conn, _ := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := proto.NewUserClient(conn)
	//data, _ := client.GetUserInfo(context.Background(), &proto.UserInfoRequest{
	//	Token: "123456",
	//})
	data, _ := client.GetUserInfo(ctx, &emptypb.Empty{})

	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
}
