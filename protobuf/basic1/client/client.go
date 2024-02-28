package main

import (
	"context"
	"encoding/json"
	"fmt"
	"goLearning/protobuf/basic1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	// 客户端调用 GetCourseList 方法
	client := proto.NewCourseClient(conn)
	data, _ := client.GetCourseList(context.Background(), &proto.CourseRequest{
		Type: "language",
	})

	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))
}
