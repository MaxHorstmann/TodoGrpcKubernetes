package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/MaxHorstmann/TodoGrpcKubernetes/pkg/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type todoserver struct {
	pb.UnimplementedTodosServer
}

func (s *todoserver) AddTodo(ctx context.Context, todo *pb.Todo) (reply *pb.AddTodoReply, err error) {
	fmt.Println(todo.Task)
	fmt.Println(todo.Done)
	fmt.Println("yes this is the new version")
	return &pb.AddTodoReply{Something: "very nice"}, nil
}

func (s *todoserver) GetAllTodos(ctx context.Context, todoParams *pb.GetAllTodosParams) (reply *pb.GetAllTodosResponse, err error) {
	return &pb.GetAllTodosResponse{Response: []string{"first one", "second one", "third one"}}, nil
}

func main() {
	fmt.Println("running")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1234))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodosServer(s, &todoserver{})
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
