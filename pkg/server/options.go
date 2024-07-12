package server

import (
	"context"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tonghia/go-challenge-transaction-app/pkg/server/gateway"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// Option configures a gRPC and a gateway server.
type Option func(*serverConfig)

func createConfig(opts []Option) *serverConfig {
	c := createDefaultConfig()
	for _, f := range opts {
		f(c)
	}
	return c
}

// WithGatewayAddr ...
func WithGatewayAddr(host string, port int) Option {
	return func(c *serverConfig) {
		c.gateway.addr = Listen{
			Host: host,
			Port: port,
		}
	}
}

// WithGatewayAddrListen ...
func WithGatewayAddrListen(l Listen) Option {
	return func(c *serverConfig) {
		c.gateway.addr = l
	}
}

// WithGatewayMuxOptions returns an Option that sets runtime.ServeMuxOption(s) to a gateway server.
func WithGatewayMuxOptions(opts ...runtime.ServeMuxOption) Option {
	return func(c *serverConfig) {
		c.gateway.muxOptions = append(c.gateway.muxOptions, opts...)
	}
}

// WithGatewayServerMiddlewares returns an Option that sets middleware(s) for http.Server to a gateway server.
func WithGatewayServerMiddlewares(middlewares ...gateway.HTTPServerMiddleware) Option {
	return func(c *serverConfig) {
		c.gateway.serverMiddlewares = append(c.gateway.serverMiddlewares, middlewares...)
	}
}

// WithGatewayServerHandler returns an Option that sets hanlers(s) for http.Server to a gateway server.
func WithGatewayServerHandler(handlers ...gateway.HTTPServerHandler) Option {
	return func(c *serverConfig) {
		c.gateway.serverHandlers = append(c.gateway.serverHandlers, handlers...)
	}
}

// WithGatewayServerConfig returns an Option that specifies http.Server configuration to a gateway server.
func WithGatewayServerConfig(cfg *gateway.HTTPServerConfig) Option {
	return func(c *serverConfig) {
		c.gateway.serverConfig = cfg
	}
}

// WithPassedHeader returns an Option that sets configurations about passed headers for a gateway server.
func WithPassedHeader(decider gateway.PassedHeaderDeciderFunc) Option {
	return WithGatewayServerMiddlewares(gateway.CreatePassingHeaderMiddleware(decider))
}

///-------------------------- GRPC options below--------------------------

// WithGrpcAddr ...
func WithGrpcAddr(host string, port int) Option {
	return func(c *serverConfig) {
		c.grpc.addr = Listen{
			Host: host,
			Port: port,
		}
	}
}

// WithGrpcAddrListen ...
func WithGrpcAddrListen(l Listen) Option {
	return func(c *serverConfig) {
		c.grpc.addr = l
	}
}

// WithGrpcServerUnaryInterceptors returns an Option that sets unary interceptor(s) for a gRPC server.
func WithGrpcServerUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(c *serverConfig) {
		c.grpc.serverUnaryInterceptors = append(c.grpc.serverUnaryInterceptors, interceptors...)
	}
}

// WithGrpcServerStreamInterceptors returns an Option that sets stream interceptor(s) for a gRPC server.
func WithGrpcServerStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(c *serverConfig) {
		c.grpc.serverStreamInterceptors = append(c.grpc.serverStreamInterceptors, interceptors...)
	}
}

// WithDefaultLogger returns an Option that sets default grpclogger.LoggerV2 object.
func WithDefaultLogger() Option {
	return func(c *serverConfig) {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))
	}
}

// WithServiceServer ...
func WithServiceServer(srv ...ServiceServer) Option {
	return func(c *serverConfig) {
		c.serviceServers = append(c.serviceServers, srv...)
	}
}

// WithTracerProvider enable tracing
func WithTracerProvider(provider trace.TracerProvider) Option {
	return func(cfg *serverConfig) {
		if provider != nil {
			cfg.grpc.tracerProvider = provider
		}
	}
}

// WithValidateCallback validate callback
func WithValidateCallback(callback func(context.Context, error)) Option {
	return func(cfg *serverConfig) {
		if callback != nil {
			cfg.grpc.validateCallback = callback
		}
	}
}
