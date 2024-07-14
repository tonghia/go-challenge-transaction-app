package pbconv

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/tonghia/go-challenge-transaction-app/internal/model"
	"github.com/tonghia/go-challenge-transaction-app/pb"
)

// Converts a valid pb.Transaction to model.AccountTransaction correctly
func TestTransactionFromPb_ValidTransaction(t *testing.T) {
	p := &pb.Transaction{
		Id:              1,
		AccountId:       2,
		Amount:          &pb.Decimal{Unit: 100, Nanos: 500000000},
		TransactionType: model.TransactionTypeDeposit,
	}

	result := TransactionFromPb(p)

	expected := &model.AccountTransaction{
		ID:              1,
		AccountID:       2,
		Amount:          decimal.New(100, 500000000),
		TransactionType: model.TransactionTypeDeposit,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Converts AccountTransaction to pb.Transaction correctly
func TestTransactionToPb_ConvertsCorrectly(t *testing.T) {
	amount := decimal.NewFromFloat(123.45)
	accountTransaction := &model.AccountTransaction{
		ID:              123,
		AccountID:       456,
		Amount:          amount,
		TransactionType: model.TransactionTypeDeposit,
	}

	expected := &pb.Transaction{
		Id:        123,
		AccountId: 456,
		Amount: &pb.Decimal{
			Unit:  amount.IntPart(),
			Nanos: amount.Exponent(),
		},
		TransactionType: model.TransactionTypeDeposit,
	}

	result := TransactionToPb(accountTransaction)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// Converts a list of AccountTransaction to a list of pb.Transaction correctly
func TestConvertsListOfAccountTransactionToListOfPbTransaction(t *testing.T) {
	transactions := []*model.AccountTransaction{
		{
			ID:     1,
			Amount: decimal.New(100, 0),
		},
		{
			ID:     2,
			Amount: decimal.New(200, 0),
		},
	}

	expected := []*pb.Transaction{
		{
			Id: 1,
			Amount: &pb.Decimal{
				Unit:  100,
				Nanos: 0,
			},
		},
		{
			Id: 2,
			Amount: &pb.Decimal{
				Unit:  200,
				Nanos: 0,
			},
		},
	}

	result := TransactionsToPb(transactions)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
