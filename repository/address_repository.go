package repository

import (
	"context"
	"database/sql"
	"time"
	"userboilerplate-api/config"
	"userboilerplate-api/entity"
)

func AddressById(ctx context.Context, id string) (entity.Address, error) {
	var result entity.Address

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).First(&result, "id = @id", sql.Named("id", id)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return result, nil
}

func AddressByUser(ctx context.Context, user_id string) ([]entity.Address, error) {
	var result []entity.Address

	config.GetConnection().Begin()
	if err := config.GetConnection().WithContext(ctx).Find(&result).Where("user_id = @user_id", sql.Named("user_id", user_id)).Error; err != nil {
		config.GetConnection().Rollback()
	}
	config.GetConnection().Commit()

	return result, nil
}

func CreateAddress(ctx context.Context, data interface{}) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Create(data).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func UpdateAddress(ctx context.Context, data entity.Address, id string) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Save(&entity.Address{Id: id, UserId: data.UserId, AddressType: data.AddressType, RecipientName: data.RecipientName, RecipientPhone: data.RecipientPhone, AddressName: data.AddressName, PostalCode: data.PostalCode, City: data.City, UpdatedAt: time.Now()}).Error; err != nil {
		config.GetConnection().Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}
