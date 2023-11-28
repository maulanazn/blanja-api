package wishlist

import (
	"config"
	"context"
	"database/sql"
)

func InsertWishlist(ctx context.Context, wishlist interface{}) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Create(wishlist).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func UpdateWishlist(ctx context.Context, data Wishlist, id string) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Model(&data).Where("id = @id", sql.Named("id", id)).Updates(&Wishlist{WishlistId: id, UserId: data.UserId, ProductId: data.ProductId, Quantity: data.Quantity}).Error; err != nil {
		config.GetConnection().Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}
