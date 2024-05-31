package grpc

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type Middleware func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)

func ExampleMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	log.Println("Example middleware")
	return handler(ctx, req)
}
