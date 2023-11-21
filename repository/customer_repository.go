package repository

import (
	"config"
	"context"
	"database/sql"
	entity "entity"
	"time"
)

func CreateCustomer(ctx context.Context, data interface{}) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Create(data).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func UpdateCustomer(ctx context.Context, data entity.Users, id string) error {
	config.GetConnection().WithContext(ctx).Begin()
	config.GetConnection().WithContext(context.Background()).Model(&data).Where("id = @id", sql.Named("id", id)).Updates(entity.Users{Id: id, Userimage: data.Userimage, Username: data.Username, Email: data.Email, Phone: data.Phone, Gender: data.Gender, Dateofbirth: data.Dateofbirth, Password: data.Password, Roles: data.Roles, UpdatedAt: time.Now()})
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func SelectEmailCustomers(ctx context.Context, email string) (entity.Users, error) {
	var result entity.Users

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).First(&result, "email = @email", sql.Named("email", email)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
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
