package repository

import (
	"context"
	"gorm.io/gorm"
	"source-base-go/golang/service/transactionService/model"
)

type TransactionRepository interface {
	SaveTransaction(ctx context.Context, info *model.Transaction) error
	SaveLog(ctx context.Context, log model.TransactionLog) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) SaveTransaction(ctx context.Context, info *model.Transaction) error {
	return r.db.WithContext(ctx).Create(&info).Error
}

func (r *transactionRepository) SaveLog(ctx context.Context, log model.TransactionLog) error {
	return r.db.WithContext(ctx).Create(&log).Error
}
