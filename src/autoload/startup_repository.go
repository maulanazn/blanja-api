package autoload

import (
	"address"
	"config"
	"log"
	"order_details"
	"users"
	"wishlist"
)

func InitDBPostgreSQL() {
	if err := config.GetConnection().AutoMigrate(&users.Users{}, &address.Address{}, &wishlist.Wishlist{}, &order_details.OrderDetail{}); err != nil {
		log.Println(err.Error())
	}
}
