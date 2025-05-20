package repository

import (
	"context"
	"gorm.io/gorm"
	"source-base-go/golang/transactionService/model"
)

type TransactionRepository interface {
	BeginTx(ctx context.Context) (*gorm.DB, error)
	CommitTx(tx *gorm.DB) error
	RollbackTx(tx *gorm.DB) error
	SaveLog(ctx context.Context, tx *gorm.DB, log model.TransactionLog) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) BeginTx(ctx context.Context) (*gorm.DB, error) {
	tx := r.db.WithContext(ctx).Begin()
	return tx, tx.Error
}

func (r *transactionRepository) CommitTx(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *transactionRepository) RollbackTx(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *transactionRepository) SaveLog(ctx context.Context, tx *gorm.DB, log model.TransactionLog) error {
	return tx.WithContext(ctx).Create(&log).Error
}
