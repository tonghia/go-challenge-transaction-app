package pbconv

import (
	"github.com/tonghia/go-challenge-transaction-app/internal/model"
	"github.com/tonghia/go-challenge-transaction-app/pb"
)

func TransactionToPb(t *model.AccountTransaction) *pb.Transaction {
	return &pb.Transaction{
		Id:              t.ID,
		AccountId:       t.AccountID,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
	}
}

func TransactionsToPb(ts []*model.AccountTransaction) []*pb.Transaction {
	rs := make([]*pb.Transaction, len(ts))
	for i, t := range ts {
		rs[i] = TransactionToPb(t)
	}

	return rs
}
