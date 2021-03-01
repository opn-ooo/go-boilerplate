package database

import (
	"context"
	"gorm.io/gorm"
)

type TransactionInterface interface {
	WithInTransaction(ctx *context.Context, fn func(ctx context.Context) error) error
}

type Transaction struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) TransactionInterface {
	return &Transaction{db}
}

var transactionKey = struct{}{}

func (r *Transaction) WithInTransaction(ctx *context.Context, fn func(ctx context.Context) error) error {
	return r.db.Transaction(func(transaction *gorm.DB) error {
		*ctx = context.WithValue(*ctx, &transactionKey, transaction)
		return fn(*ctx)
	})
}

func GetTransaction(context *context.Context) (*gorm.DB, bool) {
	if context == nil {
		return nil, false
	}
	transaction, ok := (*context).Value(&transactionKey).(*gorm.DB)
	return transaction, ok
}
