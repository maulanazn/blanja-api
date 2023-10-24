package service

import (
	"belanjabackend/config"
	"belanjabackend/entity"
	"belanjabackend/webserver/request"
	"context"
)

func CreateCustomer(ctx context.Context, req request.RegisterRequest) {
	db := config.GetConnection()

	db.Create(&entity.Customer{
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.Password,
	})
}
