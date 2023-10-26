package repository

import (
	"belanjabackend/config"
	"context"
	"database/sql"
	"errors"
)

func CreateCustomer(ctx context.Context, data interface{}) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Create(data).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
		return errors.New("Duplicate")
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func SelectEmailCustomers(ctx context.Context, data interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Table("customers").Take(&result).Where("email = @email", sql.Named("email", data.(string))).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
		return nil, errors.New("Duplicate")
	}
	config.GetConnection().WithContext(ctx).Commit()

	return result, nil
}
