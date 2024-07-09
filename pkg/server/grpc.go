package server

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	grpc_selector "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

const prefixMethodHeathCheck = "/pb.HealthService/"

var mwSelector = grpc_selector.MatchFunc(func(ctx context.Context, callMeta interceptors.CallMeta) bool {
	return !strings.HasPrefix(callMeta.FullMethod(), prefixMethodHeathCheck)
})

type grpcConfig struct {
	addr                     Listen
	serverUnaryInterceptors  []grpc.UnaryServerInterceptor
	serverStreamInterceptors []grpc.StreamServerInterceptor
	grpcOptions              []grpc.ServerOption
	tracerProvider           trace.TracerProvider
	validateCallback         func(context.Context, error)
	maxConcurrentStreams     uint32
}

func createDefaultGrpcConfig() *grpcConfig {
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.EnableClientHandlingTimeHistogram()

	config := &grpcConfig{
		addr: Listen{
			Host: "0.0.0.0",
			Port: 10000,
		},
		serverUnaryInterceptors: []grpc.UnaryServerInterceptor{
			// otelgrpc.UnaryServerInterceptor(),
			grpc_selector.UnaryServerInterceptor(
				grpc_prometheus.UnaryServerInterceptor, mwSelector),
		},
		serverStreamInterceptors: []grpc.StreamServerInterceptor{
			// otelgrpc.StreamServerInterceptor(),
			grpc_selector.StreamServerInterceptor(
				grpc_prometheus.StreamServerInterceptor, mwSelector),
		},
		maxConcurrentStreams: 1000,
	}

	return config
}

func (c *grpcConfig) appendDefaultOptions() []grpc.ServerOption {
	return append(
		[]grpc.ServerOption{
			grpc.ChainUnaryInterceptor(c.serverUnaryInterceptors...),
			grpc.ChainStreamInterceptor(c.serverStreamInterceptors...),
			grpc.MaxConcurrentStreams(c.maxConcurrentStreams),
		},
		c.grpcOptions...,
	)
}

// grpcServer wraps grpc.Server setup process.
type grpcServer struct {
	server         *grpc.Server
	config         *grpcConfig
	serviceServers []ServiceServer
}

func newGrpcServer(c *grpcConfig, servers []ServiceServer) *grpcServer {
	gs := &grpcServer{
		config:         c,
		serviceServers: servers,
	}

	validateOpts := []grpc_validator.Option{grpc_validator.WithOnValidationErrCallback(c.validateCallback)}
	gs.config.serverUnaryInterceptors =
		append(gs.config.serverUnaryInterceptors,
			grpc_validator.UnaryServerInterceptor(validateOpts...))
	gs.config.serverStreamInterceptors =
		append(gs.config.serverStreamInterceptors,
			grpc_validator.StreamServerInterceptor(validateOpts...))

	if c.tracerProvider != nil {
		otelOpts := []otelgrpc.Option{otelgrpc.WithTracerProvider(c.tracerProvider)}
		gs.config.serverUnaryInterceptors =
			append(gs.config.serverUnaryInterceptors,
				grpc_selector.UnaryServerInterceptor(
					otelgrpc.UnaryServerInterceptor(otelOpts...), mwSelector))
		gs.config.serverStreamInterceptors =
			append(gs.config.serverStreamInterceptors,
				grpc_selector.StreamServerInterceptor(
					otelgrpc.StreamServerInterceptor(otelOpts...), mwSelector))
	}

	return gs
}

// serve implements server.Server
func (s *grpcServer) serve() error {
	listener, err := s.config.addr.createListener()
	if err != nil {
		return fmt.Errorf("failed to create listener %w", err)
	}
	log.Println("gRPC server is starting", listener.Addr().String())

	s.server = grpc.NewServer(s.config.appendDefaultOptions()...)
	for _, svr := range s.serviceServers {
		svr.RegisterWithGrpcServer(s.server)
	}

	if err = s.server.Serve(listener); err != nil {
		log.Println("err while serving", err)
		return fmt.Errorf("failed to serve gRPC server %w", err)
	}
	log.Println("gRPC server ready")

	return nil
}

// shutdown ...
func (s *grpcServer) shutdown(ctx context.Context) {
	s.server.GracefulStop()
}
