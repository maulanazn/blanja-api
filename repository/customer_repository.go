package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"userboilerplate-api/config"
	entity "userboilerplate-api/entity"
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

func UpdateCustomer(ctx context.Context, data entity.Users, id string) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(context.Background()).Table("users").Where("id = @id", sql.Named("id", id)).Updates(map[string]interface{}{
		"userimage":   data.Userimage,
		"username":    data.Username,
		"roles":       data.Roles,
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
	if err := config.GetConnection().WithContext(ctx).Table("users").Take(&result).Where("email = @email", sql.Named("email", data.(string))).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
		return nil, errors.New("Duplicate")
	}
	config.GetConnection().WithContext(ctx).Commit()

	return result, nil
}

func SelectCustomerById(ctx context.Context, id string) (entity.Users, error) {
	var result entity.Users

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).First(&result, "id = @id", sql.Named("id", id)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return result, nil
}
