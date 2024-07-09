package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Server is the framework instance.
type Server struct {
	grpcServer    *grpcServer
	gatewayServer *gatewayServer
	config        *serverConfig
}

// New creates a server instance.
func New(opts ...Option) (*Server, error) {
	c := createConfig(opts)

	log.Println("create grpc server")
	grpcServer := newGrpcServer(c.grpc, c.serviceServers)

	log.Println("create gateway server")
	gatewayServer, err := newGatewayServer(c.gateway, c.grpc.addr.String(), c.serviceServers)
	if err != nil {
		return nil, fmt.Errorf("fail to create gateway server. %w", err)
	}

	return &Server{
		grpcServer:    grpcServer,
		gatewayServer: gatewayServer,
		config:        c,
	}, nil
}

// Serve starts gRPC and Gateway servers.
func (s *Server) Serve(ctx context.Context) error {
	stop := make(chan os.Signal, 1)
	errCh := make(chan error)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.gatewayServer.serve(); err != nil {
			log.Println("error starting http server", err)
			errCh <- err
		}
	}()

	go func() {
		if err := s.grpcServer.serve(); err != nil {
			log.Println("error starting gRPC server", err)
			errCh <- err
		}
	}()

	for {
		select {
		case <-stop:
			s.Stop(ctx)
			return nil

		case <-ctx.Done():
			s.Stop(ctx)
			return nil

		case err := <-errCh:
			return err
		}
	}
}

func (s *Server) Stop(ctx context.Context) {
	log.Println("Shutting down server")

	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, ss := range s.config.serviceServers {
		ss.Close(timeoutCtx)
	}

	s.gatewayServer.shutdown(ctx)
	s.grpcServer.shutdown(ctx)
}
