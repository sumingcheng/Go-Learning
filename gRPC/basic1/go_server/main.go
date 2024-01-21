package main

import (
	"context"
	"goLearning/gRPC/basic1/proto"
	"google.golang.org/grpc"
	"net"
	"time"
)

type Todo struct {
	Id        int64
	Content   string
	Completed bool
}

type TodoList struct {
	list []Todo
}

func (td *TodoList) AddTodo(ctx context.Context, req *proto.AddingTodoRequest) (*proto.TodoResponse, error) {
	todo := Todo{
		Id:        time.Now().Unix(), // long 类型
		Content:   req.Content,
		Completed: false,
	}
	td.list = append(td.list, todo)

	return &proto.TodoResponse{
		Id:        todo.Id,
		Content:   todo.Content,
		Completed: todo.Completed,
	}, nil
}

func (td *TodoList) ToggleTodo(ctx context.Context, req *proto.TogglingTodoRequest) (*proto.TodoResponse, error) {
	var target Todo

	for i := 0; i < len(td.list); i++ {
		if td.list[i].Id == req.Id {
			td.list[i].Completed = !td.list[i].Completed
			target = td.list[i]
			break
		}
	}

	return &proto.TodoResponse{
		Id:        target.Id,
		Content:   target.Content,
		Completed: target.Completed,
	}, nil
}

func (td *TodoList) RemoveTodo(ctx context.Context, req *proto.RemovingTodoRequest) (*proto.TodoResponse, error) {
	var target Todo

	for i := 0; i < len(td.list); i++ {
		if td.list[i].Id == req.Id {
			target = td.list[i]
			td.list = append(td.list[0:i], td.list[i+1:]...)
			break
		}
	}

	return &proto.TodoResponse{
		Id:        target.Id,
		Content:   target.Content,
		Completed: target.Completed,
	}, nil
}

func main() {
	// 创建服务
	gServer := grpc.NewServer()
	//注册服务
	proto.RegisterTodoListServer(gServer, &TodoList{
		list: make([]Todo, 0),
	})
	// 创建服务监听对象
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err.Error())
	}

	//使用监听对象
	err = gServer.Serve(listen)
	if err != nil {
		return
	}
}
