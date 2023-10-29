package entity

import (
	"belanjabackend/config"
	"context"
	"database/sql"
	"time"
)

type Customer struct {
	Id          string `gorm:"primaryKey"`
	Address     []Address
	Userimage   string    `json:"userimage"`
	Username    string    `json:"username"`
	Email       string    `gorm:"unique"`
	Phone       int64     `json:"phone"`
	Gender      string    `json:"gender"`
	Dateofbirth string    `json:"dateofbirth"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func (c *Customer) ValidateUpdate(id string) interface{} {
	var result map[string]interface{}
	var data interface{}

	config.GetConnection().WithContext(context.Background()).Begin()
	if err := config.GetConnection().WithContext(context.Background()).Table("customers").Take(&result).Where("id = @id", sql.Named("id", 13)).Error; err != nil {
		config.GetConnection().WithContext(context.Background()).Rollback()
		return nil
	}
	config.GetConnection().WithContext(context.Background()).Commit()

	switch {
	case c.Userimage == "":
		data = result["userimage"].(string)
		return data
	case c.Username == "":
		data = result["username"].(string)
		return data
	case c.Phone == 0:
		data = result["phone"].(int64)
		return data
	case c.Gender == "":
		data = result["gender"].(string)
		return data
	case c.Dateofbirth == "":
		data = result["dateofbirth"].(string)
		return data
	default:
		return nil
	}
}
