package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"goLearning/gRPC-advanced/basic1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type User struct {
	proto.UnimplementedUserServer
}

var Token = "123456"

func (u *User) GetUserInfo(ctx context.Context, empty *emptypb.Empty) (*proto.UserInfoResponse, error) {

	return &proto.UserInfoResponse{
		Code: 0,
		Msg:  "OK",
		Data: &proto.UserInfo{
			Uid:      1,
			Username: "user1",
			CourseList: []proto.Course{
				proto.Course_Go,
				proto.Course_Java,
			},
		},
	}, nil
}

func main() {
	opt := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			checkToken,
			Test,
		)),
	}
	gServer := grpc.NewServer(opt...)
	proto.RegisterUserServer(gServer, &User{})
	listerner, _ := net.Listen("tcp", ":8080")
	_ = gServer.Serve(listerner)
}

func checkToken(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("checkToken")
	md, _ := metadata.FromIncomingContext(ctx) // 获取metadata
	clientToken := md.Get("token")[0]          // 获取token
	if clientToken != Token {
		return &proto.UserInfoResponse{
			Code: 1,
			Msg:  "Token is invalid",
		}, nil
	}
	return handler(ctx, req)
}

func Test(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println("Test")
	return handler(ctx, req)
}
