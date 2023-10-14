package repository

import (
	"belanjabackend/entity"
	"context"
)

type WalletRepository interface {
	Insert(ctx context.Context, wallet entity.Wallet) (entity.Wallet, error)
	FindById(ctx context.Context, id string) (entity.Wallet, error)
	FindAll(ctx context.Context) ([]entity.Wallet, error)
	UpdateWallet(ctx context.Context, id string, wallet entity.Wallet) (entity.Wallet, error)
}
