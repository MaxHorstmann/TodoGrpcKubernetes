package main

import (
	"context"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	//"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"

	pb "github.com/MaxHorstmann/TodoGrpcKubernetes/pkg/todo"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"google.golang.org/grpc"
)

type todoserver struct {
	pb.UnimplementedTodosServer
}

func (s *todoserver) AddTodo(ctx context.Context, todo *pb.Todo) (reply *pb.AddTodoReply, err error) {
	fmt.Println(todo.Task)
	fmt.Println(todo.Done)
	fmt.Println("yes this is the improved new version ")

	//// https://docs.microsoft.com/en-us/azure/azure-sql/database/connect-query-go
	//
	//const connStr = "server=sql;user id=sa;password=Password1!;database=Foo"
	//db, err := sql.Open("sqlserver", connStr)
	//if err != nil {
	//	log.Fatal("Can't connect to SQL", err)
	//}
	//
	//_, err = db.Exec("insert into [Foo].dbo.Todos (Todo) VALUES (@todo)", sql.Named("todo", todo.Task))
	//if err != nil {
	//	log.Fatal("Can't query SQL", err)
	//}
	//
	//fmt.Println("SQL insert done!!!")

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
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	pb.RegisterTodosServer(s, &todoserver{})
	grpc_prometheus.Register(s)

	//reflection.Register(s)
	log.Printf("grpc server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":5678", nil)

}
