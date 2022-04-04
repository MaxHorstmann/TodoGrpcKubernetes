// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package todo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TodosClient is the client API for Todos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodosClient interface {
	AddTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*AddTodoReply, error)
}

type todosClient struct {
	cc grpc.ClientConnInterface
}

func NewTodosClient(cc grpc.ClientConnInterface) TodosClient {
	return &todosClient{cc}
}

func (c *todosClient) AddTodo(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*AddTodoReply, error) {
	out := new(AddTodoReply)
	err := c.cc.Invoke(ctx, "/todo.todos/AddTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodosServer is the server API for Todos service.
// All implementations must embed UnimplementedTodosServer
// for forward compatibility
type TodosServer interface {
	AddTodo(context.Context, *Todo) (*AddTodoReply, error)
	mustEmbedUnimplementedTodosServer()
}

// UnimplementedTodosServer must be embedded to have forward compatible implementations.
type UnimplementedTodosServer struct {
}

func (UnimplementedTodosServer) AddTodo(context.Context, *Todo) (*AddTodoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTodo not implemented")
}
func (UnimplementedTodosServer) mustEmbedUnimplementedTodosServer() {}

// UnsafeTodosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodosServer will
// result in compilation errors.
type UnsafeTodosServer interface {
	mustEmbedUnimplementedTodosServer()
}

func RegisterTodosServer(s grpc.ServiceRegistrar, srv TodosServer) {
	s.RegisterService(&Todos_ServiceDesc, srv)
}

func _Todos_AddTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).AddTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.todos/AddTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).AddTodo(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

// Todos_ServiceDesc is the grpc.ServiceDesc for Todos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.todos",
	HandlerType: (*TodosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTodo",
			Handler:    _Todos_AddTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo/services.proto",
}
