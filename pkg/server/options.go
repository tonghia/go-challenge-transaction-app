package server

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tonghia/go-challenge-transaction-app/pkg/server/gateway"
	"google.golang.org/grpc"
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

// WithServiceServer ...
func WithServiceServer(srv ...ServiceServer) Option {
	return func(c *serverConfig) {
		c.serviceServers = append(c.serviceServers, srv...)
	}
}
