package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"goLearning/gRPC-advanced/basic3/proto"
	"google.golang.org/grpc"
)

type User struct {
	Username string
	Password string
	proto.UnimplementedUserServer
}

const SECRET_KEY = "sjadkasjkldjaskld"

func (u *User) Login(
	ctx context.Context,
	req *proto.UserLoginRequest,
) (*proto.UserLoginResponse, error) {
	if req.Username == "sumingcheng" && req.Password == "123456" {
		claims := &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   req.Username,
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, _ := token.SignedString([]byte(SECRET_KEY))

		return &proto.UserLoginResponse{
			Code:    0,
			Message: "success",
			Token:   tokenString,
		}, nil
	}

	return &proto.UserLoginResponse{
		Code:    1,
		Message: "username or password error",
	}, nil
}

func createAccessLogInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp any, err error) {
	reqLog := "请求日志: 方法: %s , 开始时间: %d , 请求: %v"
	beginTime := time.Now().Local().Unix()
	log.Printf(reqLog, info.FullMethod, beginTime, req)

	resp, err = handler(ctx, req)

	respLog := "响应日志: 方法: %s , 结束时间: %d , 响应: %v"
	endTime := time.Now().Local().Unix()
	log.Printf(respLog, info.FullMethod, endTime, resp)

	return resp, err
}

func main() {
	opt := grpc.UnaryInterceptor(createAccessLogInterceptor)
	gServer := grpc.NewServer(opt)
	proto.RegisterUserServer(gServer, &User{})

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server is running on port :8080")
	if err := gServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
