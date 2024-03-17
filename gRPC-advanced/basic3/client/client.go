package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"goLearning/gRPC-advanced/basic3/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(clientTimeout()),
		),
	}

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewUserClient(conn)
	data, err := client.Login(context.Background(), &proto.UserLoginRequest{
		Username: "sumingcheng",
		Password: "123456",
	})
	if err != nil {
		log.Fatalf("login failed: %v", err)
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("failed to marshal json: %v", err)
	}
	fmt.Println(string(jsonData))
}

func clientTimeout() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx, cancel := createDefaultTimeout(ctx)
		if cancel != nil {
			defer cancel()
		}

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func createDefaultTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	var cancel context.CancelFunc
	if _, ok := ctx.Deadline(); !ok {
		ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	} else {
		cancel = func() {}
	}

	return ctx, cancel
}
