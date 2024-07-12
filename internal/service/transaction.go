package service

import (
	"context"
	"fmt"

	"github.com/tonghia/go-challenge-transaction-app/pb"
	"github.com/tonghia/go-challenge-transaction-app/pkg/pbconv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) ListTransactions(ctx context.Context, req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	transactions, err := s.repos.AccountTransaction.GetByUserAccount(ctx, req.UserId, req.AccountId)
	if err != nil {
		s.log.Error(fmt.Sprintf("failed to list transactions: %v", err))
		return nil, status.Error(codes.Internal, "failed to list transactions")
	}

	return &pb.ListTransactionsResponse{
		Message: codes.OK.String(),
		Data: &pb.ListTransactionsResponse_Data{
			Transactions: pbconv.TransactionsToPb(transactions),
		}}, nil
}

func (s *service) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {

	return &pb.CreateTransactionResponse{Message: codes.OK.String()}, nil
}

func (s *service) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {

	return &pb.UpdateTransactionResponse{Message: codes.OK.String()}, nil
}

func (s *service) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error) {

	return &pb.DeleteTransactionResponse{Message: codes.OK.String()}, nil
}
