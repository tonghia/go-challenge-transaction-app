package app

import (
	"context"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	grpcpv "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"github.com/tonghia/go-challenge-transaction-app/internal/config"
	"github.com/tonghia/go-challenge-transaction-app/internal/must"
	"github.com/tonghia/go-challenge-transaction-app/internal/repository"
	"github.com/tonghia/go-challenge-transaction-app/internal/service"
	"github.com/tonghia/go-challenge-transaction-app/pkg/logger"
	"github.com/tonghia/go-challenge-transaction-app/pkg/server"
)

func Run(_ []string) error {
	var ctx = context.TODO()
	var ll = logger.NewLogger()

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	db, err := must.ConnectMySQL(cfg.MySQL)
	if err != nil {
		return fmt.Errorf("failed to connect mysql: %v", err)
	}
	defer must.Close(db)
	repo := repository.NewRepository(db)

	pv, err := protovalidate.New()
	if err != nil {
		return fmt.Errorf("failed to initialize validator: %v", err)
	}

	mw := NewCustomMiddleware(ll)
	srv, err := server.New(
		server.WithGrpcAddrListen(cfg.Server.GRPC),
		server.WithGatewayAddrListen(cfg.Server.HTTP),
		server.WithServiceServer(
			service.NewService(ctx, cfg, ll, repo),
		),
		server.WithGatewayServerMiddlewares(mw.MiddlewareHandleFunc),
		server.WithGrpcServerUnaryInterceptors(grpcpv.UnaryServerInterceptor(pv)),
	)
	if err != nil {
		return fmt.Errorf("initialize server %v", err)
	}

	if err := srv.Serve(ctx); err != nil {
		return fmt.Errorf("serving %v", err)
	}

	return nil
}
