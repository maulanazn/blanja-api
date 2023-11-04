package entity

import (
	"time"
)

type AccountNumber struct {
	AccNumber      string    `gorm:"type:varchar;unique;notNull;primaryKey;column:acc_number"`
	AccOwner       string    `gorm:"type:varchar;column:acc_owner"`
	AccDateCreated time.Time `gorm:"type:date;autoCreateTime;notNull;column:acc_date_created;default:now()"`
	AccBalance     int       `gorm:"type:decimal(10,0);column:acc_balance;notNull"`
}
