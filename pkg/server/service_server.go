package server

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// ServiceServer ...
type ServiceServer interface {
	RegisterWithGrpcServer(*grpc.Server)
	RegisterWithMuxServer(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
	Close(context.Context)
}
