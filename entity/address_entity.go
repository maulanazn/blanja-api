package entity

import (
	"belanjabackend/config"
	"context"
	"database/sql"
	"time"
)

type Address struct {
	Id             string `gorm:"primaryKey"`
	CustomerId     string
	AddressType    string
	RecipientName  string
	RecipientPhone int
	AddressName    string
	PostalCode     string
	City           string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

func (adr *Address) ValidateUpdate(id string) interface{} {
	var result map[string]interface{}
	var data interface{}

	config.GetConnection().WithContext(context.Background()).Begin()
	if err := config.GetConnection().WithContext(context.Background()).Table("addresses").Take(&result).Where("id = @id", sql.Named("id", id)).Error; err != nil {
		config.GetConnection().WithContext(context.Background()).Rollback()
		return nil
	}
	config.GetConnection().WithContext(context.Background()).Commit()

	switch {
	case adr.AddressType == "":
		data = result["address_type"].(string)
		return data
	case adr.RecipientName == "":
		data = result["recipient_name"].(string)
		return data
	case adr.RecipientPhone == 0:
		data = result["recipient_phone"].(string)
		return data
	case adr.AddressName == "":
		data = result["address_name"].(string)
		return data
	case adr.PostalCode == "":
		data = result["postal_code"].(string)
		return data
	case adr.City == "":
		data = result["city"].(string)
		return data
	default:
		return nil
	}
}
