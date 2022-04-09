package main

import (
	"context"
	"log"
	"time"

	pb "github.com/MaxHorstmann/TodoGrpcKubernetes/pkg/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTodosClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.AddTodo(ctx, &pb.Todo{Done: false, Task: "Clean up kitchen"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

}
