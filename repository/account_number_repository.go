package repository

import (
	"config"
	"context"
)

func InsertAccountNumber(ctx context.Context, data interface{}) error {
	config.GetConnection().Begin()
	if err := config.GetConnection().WithContext(ctx).Create(data).Error; err != nil {
		config.GetConnection().Rollback()
	}
	config.GetConnection().Commit()

	return nil
}
