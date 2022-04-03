package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/MaxHorstmann/TodoGrpcKubernetes/todo"
	"google.golang.org/grpc"
)

type todoserver struct {
	pb.UnimplementedTodosServer
}

func main() {
	fmt.Println("running")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1234))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTodosServer(s, &todoserver{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
