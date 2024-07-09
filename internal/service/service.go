package service

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tonghia/go-challenge-transaction-app/internal/config"
	"github.com/tonghia/go-challenge-transaction-app/internal/repository"
	"github.com/tonghia/go-challenge-transaction-app/pb"
	"github.com/tonghia/go-challenge-transaction-app/pkg/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	Version = "1.0.0"
)

type serviceInterface interface {
}

var _ serviceInterface = &service{}

type service struct {
	cfg   *config.Config
	log   *zap.SugaredLogger
	repos *repository.Repository
}

func NewService(
	ctx context.Context,
	cfg *config.Config,
	log *zap.SugaredLogger,
	repos *repository.Repository,
) server.ServiceServer {

	return &service{
		cfg:   cfg,
		log:   log,
		repos: repos,
	}
}

func (s *service) RegisterWithGrpcServer(server *grpc.Server) {
	pb.RegisterServiceServer(server, s)
}

func (s *service) RegisterWithMuxServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	if err := pb.RegisterServiceHandler(ctx, mux, conn); err != nil {
		return err
	}

	return nil
}

// Close ...
func (s *service) Close(ctx context.Context) {
	s.repos.Close()
}
