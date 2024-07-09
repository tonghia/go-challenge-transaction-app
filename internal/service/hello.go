package service

import (
	"context"

	"github.com/tonghia/go-challenge-transaction-app/pb"
)

func (s *service) Hello(_ context.Context, _ *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{}, nil
}
