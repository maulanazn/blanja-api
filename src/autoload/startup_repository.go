package autoload

import (
	"address"
	"attribute"
	"config"
	"log"
	"order_details"
	"product"
	"users"
	"wishlist"
)

func InitDBPostgreSQL() {
	if err := config.GetConnection().AutoMigrate(&users.Users{}, &address.Address{}, &product.Products{}, &attribute.Attribute{}, &wishlist.Wishlist{}, &order_details.OrderDetail{}); err != nil {
		log.Println(err.Error())
	}
}
