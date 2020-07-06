package grpc

import (
	"google.golang.org/grpc"
	"github.com/AlexSwiss/todo-microservice/pkg/api/v1"
)

// RunServer runs gRPC service to publish toto service
func RunServer (ctx context.Context, v1API v1.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
}

// Register service
server := grpc.NewServer()
v1.RegisterToDoServiceServer(server, v1API)

// Graceful shutdown