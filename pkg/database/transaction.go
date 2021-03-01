package database

import (
	"context"

	"gorm.io/gorm"
)

type TransactionInterface interface {
	WithInTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type transactionManager struct {
	db *gorm.DB
}

func NewTransaction(db *gorm.DB) TransactionInterface {
	return &transactionManager{db}
}

var transactionKey = struct{}{}

func (r *transactionManager) WithInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, &transactionKey, tx)
		return fn(ctx)
	})
}

func GetTransaction(context context.Context) (*gorm.DB, bool) {
	transaction, ok := context.Value(&transactionKey).(*gorm.DB)
	return transaction, ok
}
