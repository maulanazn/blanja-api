package repository

import (
	"belanjabackend/config"
	"belanjabackend/entity"
	"context"
	"database/sql"
	"errors"
	"time"
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

func UpdateCustomer(ctx context.Context, data entity.Customer, id int64) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(context.Background()).Table("customers").Where("id = @id", sql.Named("id", id)).Updates(map[string]interface{}{
		"userimage":   data.Userimage,
		"username":    data.Username,
		"phone":       data.Phone,
		"gender":      data.Gender,
		"dateofbirth": data.Dateofbirth,
		"updated_at":  time.Now(),
	}).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
		return errors.New("failed to update")
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

func SelectCustomerById(ctx context.Context, id int64) (map[string]interface{}, error) {
	var result map[string]interface{}

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Table("customers").Take(&result).Where("id = @id", sql.Named("id", 13)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
		return nil, errors.New("Duplicate")
	}
	config.GetConnection().WithContext(ctx).Commit()

	return result, nil
}
