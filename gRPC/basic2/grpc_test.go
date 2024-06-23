package main

import (
	"context"
	"fmt"
	student_service "goLearning/gRPC/basic2/idl/my_proto"
	"google.golang.org/grpc"
	"testing"
)

func TestService(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接grpc服务端失败: %v\n", err)
		t.Fail()
	}
	defer conn.Close()
	clinet := student_service.NewStudentServiceClient(conn)

	rep, err := clinet.GetStudentInfo(context.Background(), &student_service.Request{StudentId: "学生1"})
	if err != nil {
		fmt.Printf("grpc 请求失败: %v\n", err)
		t.Fail()
	}
	fmt.Printf("调用成功！ Name %s Age %d Height %.1f\n", rep.Name, rep.Age, rep.Height)
}
