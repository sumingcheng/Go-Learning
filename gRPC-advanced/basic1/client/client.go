package main

import (
	"context"
	"encoding/json"
	"fmt"
	"goLearning/gRPC-advanced/basic1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := proto.NewUserClient(conn)
	data, _ := client.GetUserInfo(context.Background(), &proto.UserInfoRequest{
		Token: "123456",
	})

	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
}
