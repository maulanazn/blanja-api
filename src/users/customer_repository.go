package users

import (
	"config"
	"context"
	"database/sql"
	"time"
)

func InsertCustomer(ctx context.Context, data interface{}) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Create(data).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func UpdateCustomer(ctx context.Context, data Users, id string) error {
	config.GetConnection().WithContext(ctx).Begin()
	config.GetConnection().WithContext(context.Background()).Model(&data).Where("id = @id", sql.Named("id", id)).Updates(Users{Id: id, UserImage: data.UserImage, Username: data.Username, Email: data.Email, Phone: data.Phone, Gender: data.Gender, DateOfBirth: data.DateOfBirth, Password: data.Password, Roles: data.Roles, UpdatedAt: time.Now()})
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func SelectEmailCustomers(ctx context.Context, email string) (Users, error) {
	var result Users

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).First(&result, "email = @email", sql.Named("email", email)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return result, nil
}

func SelectCustomerById(ctx context.Context, id string) (Users, error) {
	var result Users

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).First(&result, "id = @id", sql.Named("id", id)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return result, nil
}
