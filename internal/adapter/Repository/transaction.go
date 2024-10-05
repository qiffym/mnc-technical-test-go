package Repository

import (
	"context"
	"mncPaymentAPI/internal/domains"

	"gorm.io/gorm"
)

type (
	TrxRepo struct {
		db *gorm.DB
	}

	TrxRepoInterface interface {
		SendTrx(
			ctx context.Context,
			trx *domains.Transaction,
		) (*domains.Transaction, error)

		GetLogTrx(
			ctx context.Context,
		) ([]domains.Transaction, error)
	}
)

func NewTrxRepo(db *gorm.DB) TrxRepoInterface {
	return TrxRepo{db: db}
}

func (repo TrxRepo) SendTrx(
	ctx context.Context,
	trx *domains.Transaction,
) (*domains.Transaction, error) {
	err := repo.db.WithContext(ctx).Create(&trx).Error
	return trx, err
}

func (repo TrxRepo) GetLogTrx(
	ctx context.Context,
) ([]domains.Transaction, error) {
	var trx []domains.Transaction
	err := repo.db.WithContext(ctx).Find((&trx)).Error
	return trx, err
}
