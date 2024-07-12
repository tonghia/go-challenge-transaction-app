package pbconv

import (
	"github.com/shopspring/decimal"
	"github.com/tonghia/go-challenge-transaction-app/internal/model"
	"github.com/tonghia/go-challenge-transaction-app/pb"
)

func TransactionToPb(t *model.AccountTransaction) *pb.Transaction {
	return &pb.Transaction{
		Id:        t.ID,
		AccountId: t.AccountID,
		Amount: &pb.Decimal{
			Unit:  t.Amount.IntPart(),
			Nanos: t.Amount.Exponent(),
		},
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

func TransactionFromPb(t *pb.Transaction) *model.AccountTransaction {
	return &model.AccountTransaction{
		ID:              t.Id,
		AccountID:       t.AccountId,
		Amount:          decimal.New(t.Amount.Unit, t.Amount.Nanos),
		TransactionType: t.TransactionType,
	}
}
