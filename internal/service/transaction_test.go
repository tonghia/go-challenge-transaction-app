package service

import (
	"context"
	"errors"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/tonghia/go-challenge-transaction-app/internal/model"
	"github.com/tonghia/go-challenge-transaction-app/internal/repository"
	"github.com/tonghia/go-challenge-transaction-app/internal/repository/mock"
	"github.com/tonghia/go-challenge-transaction-app/pb"
	"github.com/tonghia/go-challenge-transaction-app/pkg/logger"
	"github.com/tonghia/go-challenge-transaction-app/pkg/pbconv"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func TestListTransactions_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccTxnRepo := mock.NewMockAccountTransactionRepositorier(ctrl)
	s := &service{
		repos: &repository.Repository{
			AccountTransaction: mockAccTxnRepo,
		},
		log: &zap.SugaredLogger{},
	}

	txn1 := &model.AccountTransaction{
		ID:              1,
		AccountID:       1,
		Amount:          decimal.New(100, 0),
		TransactionType: model.TransactionTypeDeposit,
	}
	txn2 := &model.AccountTransaction{
		ID:              1,
		AccountID:       2,
		Amount:          decimal.New(100, 0),
		TransactionType: model.TransactionTypeDeposit,
	}

	tests := []struct {
		name      string
		req       *pb.ListTransactionsRequest
		mockQuery []*model.AccountTransaction
		want      *pb.ListTransactionsResponse
	}{
		{
			name: "Valid request with userID and accountID",
			req: &pb.ListTransactionsRequest{
				UserId:    1,
				AccountId: 1,
			},
			mockQuery: []*model.AccountTransaction{
				txn1,
			},
			want: &pb.ListTransactionsResponse{
				Code:    int32(codes.OK),
				Message: codes.OK.String(),
				Data: &pb.ListTransactionsResponse_Data{
					Transactions: []*pb.Transaction{
						pbconv.TransactionToPb(txn1),
					},
				},
			},
		},
		{
			name: "Valid request with userID and zero accountID",
			req: &pb.ListTransactionsRequest{
				UserId:    1,
				AccountId: 0,
			},
			mockQuery: []*model.AccountTransaction{
				txn1,
				txn2,
			},
			want: &pb.ListTransactionsResponse{
				Code:    int32(codes.OK),
				Message: codes.OK.String(),
				Data: &pb.ListTransactionsResponse_Data{
					Transactions: []*pb.Transaction{
						pbconv.TransactionToPb(txn1),
						pbconv.TransactionToPb(txn2),
					},
				},
			},
		},
	}

	mockAccTxnRepo.EXPECT().
		GetByUserAccount(gomock.Any(), int64(1), int64(1)).
		Return([]*model.AccountTransaction{
			txn1,
		}, nil).Times(1)
	mockAccTxnRepo.EXPECT().
		GetByUser(gomock.Any(), int64(1)).
		Return([]*model.AccountTransaction{
			txn1,
			txn2,
		}, nil).Times(1)

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			got, err := s.ListTransactions(context.Background(), test.req)

			assert.Equal(t, got, test.want)
			assert.NoError(t, err)
		})
	}
}

func TestListTransactions_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAccTxnRepo := mock.NewMockAccountTransactionRepositorier(ctrl)
	s := &service{
		repos: &repository.Repository{
			AccountTransaction: mockAccTxnRepo,
		},
		log: logger.NewLogger(),
	}

	t.Run("error from query transactions with account id", func(t *testing.T) {
		req := &pb.ListTransactionsRequest{
			UserId:    1,
			AccountId: 1,
		}

		mockAccTxnRepo.EXPECT().
			GetByUserAccount(gomock.Any(), gomock.Any(), gomock.Any()).
			Return(nil, errors.New("internal")).Times(1)

		_, err := s.ListTransactions(context.Background(), req)

		assert.ErrorIs(t, err, status.Error(codes.Internal, "failed to list transactions"))
	})

	t.Run("error from query transactions without account id", func(t *testing.T) {
		req := &pb.ListTransactionsRequest{
			UserId:    1,
			AccountId: 0,
		}

		mockAccTxnRepo.EXPECT().
			GetByUser(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("internal")).Times(1)

		_, err := s.ListTransactions(context.Background(), req)

		assert.ErrorIs(t, err, status.Error(codes.Internal, "failed to list transactions"))
	})
}

func TestCreateTransaction_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	// Test case: successful creation
	ctx := context.Background()
	req := &pb.CreateTransactionRequest{
		Transaction: &pb.CreateTransaction{
			AccountId: 1,
			Amount: &pb.Decimal{
				Unit:  100,
				Nanos: 0,
			},
			TransactionType: "deposit",
		},
	}
	mockAccountRepo := mock.NewMockAccountRepositorier(ctrl)
	mockAccountTransactionRepo := mock.NewMockAccountTransactionRepositorier(ctrl)
	s := &service{
		repos: &repository.Repository{
			Account:            mockAccountRepo,
			AccountTransaction: mockAccountTransactionRepo,
		},
		log: logger.NewLogger(),
	}

	mockAccountRepo.EXPECT().
		GetByID(gomock.Any(), gomock.Any()).
		Return(&model.Account{}, nil).Times(1)
	mockAccountTransactionRepo.EXPECT().
		CreateOne(gomock.Any(), gomock.Any()).
		Return(nil).Times(1)
	resp, err := s.CreateTransaction(ctx, req)
	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, codes.OK, status.Code(err))
	assert.Equal(t, req.Transaction, resp.Data.Transaction)
}

func TestCreateTransaction_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	req := &pb.CreateTransactionRequest{
		Transaction: &pb.CreateTransaction{AccountId: 1, Amount: &pb.Decimal{Unit: 100, Nanos: 0}, TransactionType: "deposit"},
	}

	mockAccountRepo := mock.NewMockAccountRepositorier(ctrl)
	mockAccountTransactionRepo := mock.NewMockAccountTransactionRepositorier(ctrl)
	s := &service{
		repos: &repository.Repository{
			Account:            mockAccountRepo,
			AccountTransaction: mockAccountTransactionRepo,
		},
		log: logger.NewLogger(),
	}

	t.Run("error account not found", func(t *testing.T) {
		mockAccountRepo.EXPECT().
			GetByID(gomock.Any(), gomock.Any()).
			Return(nil, gorm.ErrRecordNotFound).Times(1)
		resp, err := s.CreateTransaction(ctx, req)
		assert.Nil(t, resp)
		assert.Equal(t, codes.NotFound, status.Code(err))
		assert.Equal(t, "account not found", status.Convert(err).Message())
	})

	t.Run("error failed to get account", func(t *testing.T) {

		mockAccountRepo.EXPECT().
			GetByID(gomock.Any(), gomock.Any()).
			Return(nil, errors.New("internal")).Times(1)
		resp, err := s.CreateTransaction(ctx, req)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Equal(t, "failed to get account", status.Convert(err).Message())
	})

	t.Run("error failed to create transaction", func(t *testing.T) {

		mockAccountRepo.EXPECT().
			GetByID(gomock.Any(), gomock.Any()).
			Return(new(model.Account), nil).Times(1)
		mockAccountTransactionRepo.EXPECT().
			CreateOne(gomock.Any(), gomock.Any()).
			Return(errors.New("internal")).Times(1)
		resp, err := s.CreateTransaction(ctx, req)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Equal(t, "failed to create transaction", status.Convert(err).Message())
	})
}

func TestUpdateTransaction_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	req := &pb.UpdateTransactionRequest{
		TransactionId: 1,
		Transaction: &pb.Transaction{
			AccountId: 1,
			Amount: &pb.Decimal{
				Unit:  100,
				Nanos: 0,
			},
		},
	}

	mockAccountRepo := mock.NewMockAccountRepositorier(ctrl)
	mockAccountTransactionRepo := mock.NewMockAccountTransactionRepositorier(ctrl)
	mockLogger := logger.NewLogger()

	s := &service{
		repos: &repository.Repository{
			Account:            mockAccountRepo,
			AccountTransaction: mockAccountTransactionRepo,
		},
		log: mockLogger,
	}

	mockAccountRepo.EXPECT().
		GetByID(gomock.Any(), int64(1)).
		Return(&model.Account{}, nil).Times(1)
	mockAccountTransactionRepo.EXPECT().
		UpdateOne(gomock.Any(), gomock.Any()).
		Return(nil).Times(1)

	resp, err := s.UpdateTransaction(ctx, req)

	assert.NotNil(t, resp)
	assert.Nil(t, err)
	assert.Equal(t, codes.OK, status.Code(err))
	assert.Equal(t, req.Transaction, resp.Data.Transaction)
}

func TestUpdateTransaction_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	req := &pb.UpdateTransactionRequest{
		TransactionId: 1,
		Transaction: &pb.Transaction{
			AccountId: 1,
			Amount: &pb.Decimal{
				Unit:  100,
				Nanos: 0,
			},
		},
	}

	mockAccountRepo := mock.NewMockAccountRepositorier(ctrl)
	mockAccountTransactionRepo := mock.NewMockAccountTransactionRepositorier(ctrl)
	mockLogger := logger.NewLogger()

	s := &service{
		repos: &repository.Repository{
			Account:            mockAccountRepo,
			AccountTransaction: mockAccountTransactionRepo,
		},
		log: mockLogger,
	}

	t.Run("account not found", func(t *testing.T) {
		mockAccountRepo.EXPECT().
			GetByID(gomock.Any(), int64(1)).
			Return(nil, gorm.ErrRecordNotFound).Times(1)
		resp, err := s.UpdateTransaction(ctx, req)
		assert.Nil(t, resp)
		assert.Equal(t, codes.NotFound, status.Code(err))
		assert.Equal(t, "account not found", status.Convert(err).Message())
	})

	t.Run("failed to get account", func(t *testing.T) {
		mockAccountRepo.EXPECT().
			GetByID(gomock.Any(), int64(1)).
			Return(nil, errors.New("internal")).Times(1)
		resp, err := s.UpdateTransaction(ctx, req)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Equal(t, "failed to get account", status.Convert(err).Message())
	})

	t.Run("failed to update transaction", func(t *testing.T) {
		mockAccountRepo.EXPECT().
			GetByID(gomock.Any(), int64(1)).
			Return(&model.Account{}, nil).Times(1)
		mockAccountTransactionRepo.EXPECT().
			UpdateOne(gomock.Any(), gomock.Any()).
			Return(errors.New("internal")).Times(1)
		resp, err := s.UpdateTransaction(ctx, req)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Equal(t, "failed to update transaction", status.Convert(err).Message())
	})
}

func TestDeleteTransaction_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	req := &pb.DeleteTransactionRequest{
		TransactionId: 1,
	}

	mockAccountTransactionRepo := mock.NewMockAccountTransactionRepositorier(ctrl)
	mockLogger := logger.NewLogger()

	s := &service{
		repos: &repository.Repository{
			AccountTransaction: mockAccountTransactionRepo,
		},
		log: mockLogger,
	}

	mockAccountTransactionRepo.EXPECT().
		DeleteByTransactionID(gomock.Any(), int64(1)).
		Return(nil).Times(1)

	resp, err := s.DeleteTransaction(ctx, req)

	assert.NotNil(t, resp)
	assert.Equal(t, codes.OK, status.Code(err))
	assert.Nil(t, err)
}

func TestDeleteTransaction_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	req := &pb.DeleteTransactionRequest{
		TransactionId: 1,
	}

	mockAccountTransactionRepo := mock.NewMockAccountTransactionRepositorier(ctrl)
	mockLogger := logger.NewLogger()

	s := &service{
		repos: &repository.Repository{
			AccountTransaction: mockAccountTransactionRepo,
		},
		log: mockLogger,
	}

	t.Run("failed to delete transaction", func(t *testing.T) {
		mockAccountTransactionRepo.EXPECT().
			DeleteByTransactionID(gomock.Any(), int64(1)).
			Return(errors.New("internal")).Times(1)
		resp, err := s.DeleteTransaction(ctx, req)
		assert.Nil(t, resp)
		assert.Equal(t, codes.Internal, status.Code(err))
		assert.Equal(t, "failed to delete transaction", status.Convert(err).Message())
	})
}
