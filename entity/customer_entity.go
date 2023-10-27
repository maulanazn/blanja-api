package entity

import (
	"belanjabackend/repository"
	"belanjabackend/webserver/helper"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Userimage   interface{} `json:"userimage"`
	Username    string      `json:"username"`
	Email       string      `gorm:"unique"`
	Phone       int64       `json:"phone"`
	Gender      string      `json:"gender"`
	Dateofbirth string      `json:"dateofbirth"`
	Password    string      `json:"password"`
	CreatedAt   time.Time   `json:"createdat"`
	UpdatedAt   time.Time   `json:"updatedat"`
	DeletedAt   time.Time   `json:"deletedat"`
}

func (c *Customer) ValidateUpdate(id int64) (interface{}, error) {
	var data interface{}
	result, resultErr := repository.SelectCustomerById(context.Background(), id)
	helper.PanicIfError(resultErr)

	switch {
	case c.Userimage.(string) == "":
		data = result["userimage"].(string)
		return data, nil
	case c.Username == "":
		data = result["username"].(string)
		return data, nil
	case c.Phone == 0:
		data = result["phone"].(int64)
		return data, nil
	case c.Gender == "":
		data = result["gender"].(string)
		return data, nil
	case c.Dateofbirth == "":
		data = result["dateofbirth"].(string)
		return data, nil
	default:
		return nil, errors.New("what!?")
	}
}
