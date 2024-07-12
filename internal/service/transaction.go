package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/tonghia/go-challenge-transaction-app/pb"
	"github.com/tonghia/go-challenge-transaction-app/pkg/pbconv"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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
	if _, err := s.repos.Account.GetByID(ctx, req.Transaction.AccountId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "account not found")
		} else {
			s.log.Error(fmt.Sprintf("failed to get account: %v", err))
			return nil, status.Error(codes.Internal, "failed to get account")
		}
	}

	newTransaction := pbconv.TransactionFromPb(req.Transaction)
	if err := s.repos.AccountTransaction.CreateOne(ctx, newTransaction); err != nil {
		s.log.Error(fmt.Sprintf("failed to create transaction: %v", err))
		return nil, status.Error(codes.Internal, "failed to create transaction")
	}

	return &pb.CreateTransactionResponse{
		Message: codes.OK.String(),
		Data:    &pb.CreateTransactionResponse_Data{Transaction: req.Transaction},
	}, nil
}

func (s *service) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {
	if _, err := s.repos.Account.GetByID(ctx, req.TransactionId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "account not found")
		} else {
			s.log.Error(fmt.Sprintf("failed to get account: %v", err))
			return nil, status.Error(codes.Internal, "failed to get account")
		}
	}

	updateTransaction := pbconv.TransactionFromPb(req.Transaction)
	if err := s.repos.AccountTransaction.UpdateOne(ctx, updateTransaction); err != nil {
		s.log.Error(fmt.Sprintf("failed to update transaction: %v", err))
		return nil, status.Error(codes.Internal, "failed to update transaction")
	}

	return &pb.UpdateTransactionResponse{
		Message: codes.OK.String(),
		Data:    &pb.UpdateTransactionResponse_Data{Transaction: req.Transaction},
	}, nil
}

func (s *service) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error) {
	if _, err := s.repos.Account.GetByID(ctx, req.TransactionId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "account not found")
		} else {
			s.log.Error(fmt.Sprintf("failed to get account: %v", err))
			return nil, status.Error(codes.Internal, "failed to get account")
		}
	}

	if err := s.repos.AccountTransaction.DeleteByTransactionID(ctx, req.TransactionId); err != nil {
		s.log.Error(fmt.Sprintf("failed to delete transaction: %v", err))
		return nil, status.Error(codes.Internal, "failed to delete transaction")
	}

	return &pb.DeleteTransactionResponse{Message: codes.OK.String()}, nil
}
