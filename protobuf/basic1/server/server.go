package main

import (
	"context"
	"goLearning/protobuf/basic1/proto"
	"goLearning/protobuf/basic3/data"
	"google.golang.org/grpc"
	"net"
)

type Course struct {
	proto.UnimplementedCourseServer // 这是 gRPC 最新版本推荐的实践，确保如果有未实现的方法，编译器可以给出警告。
}

func (c *Course) GetCourseList(ctx context.Context, req *proto.CourseRequest) (*proto.CourseResponse, error) {
	switch req.Type {
	case "language":
		return &proto.CourseResponse{
			Code: 0,
			Msg:  "success",
			Data: data.LanguageCourses,
			Extra: map[string]string{
				"test": "test",
			},
		}, nil
	case "computer":
		return &proto.CourseResponse{
			Code: 0,
			Msg:  "success",
			Data: data.ComputerCourses,
		}, nil
	default:
		return &proto.CourseResponse{
			Code: 1,
			Msg:  "error",
		}, nil
	}
}

func main() {
	gServer := grpc.NewServer()
	proto.RegisterCourseServer(gServer, &Course{})
	listener, _ := net.Listen("tcp", ":8080")
	_ = gServer.Serve(listener)
}
