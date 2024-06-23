package main

import (
	"context"
	"errors"
	"fmt"
	"goLearning/gRPC/basic2/idl/my_proto"
	"google.golang.org/grpc"
	"net"
)

type StudentServer struct {
	student_service.UnimplementedStudentServiceServer
}

func (s *StudentServer) GetStudentInfo(
	ctx context.Context,
	request *student_service.Request,
) (*student_service.Student, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("接口执行出错：%v\n", err)
		}
	}()

	studentId := request.StudentId
	if len(studentId) == 0 {
		return nil, errors.New("studentId is empty")
	}
	student := &student_service.Student{
		Name:      "sumingcheng",
		Age:       0,
		Gender:    false,
		Height:    0,
		Locations: nil,
		Scores:    nil,
		Birthday:  nil,
	}

	return student, nil
}

func main() {
	list, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("服务运行在:127.0.0.1:8888")
	server := grpc.NewServer()
	student_service.RegisterStudentServiceServer(server, new(StudentServer))
	err = server.Serve(list)
	if err != nil {
		panic(err)
	}
}
