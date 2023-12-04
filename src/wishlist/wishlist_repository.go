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

func SelectWishlistByStore(ctx context.Context, store_name string) ([]Wishlist, error) {
	var result []Wishlist

	config.GetConnection().Begin()
	if err := config.GetConnection().WithContext(ctx).Find(&result).Where("store_name = @store_name", sql.Named("store_name", store_name)).Error; err != nil {
		config.GetConnection().Rollback()
	}
	config.GetConnection().Commit()

	return result, nil
}

func SelectWishlistByUser(ctx context.Context, user_id string) ([]Wishlist, error) {
	var result []Wishlist

	config.GetConnection().Begin()
	if err := config.GetConnection().WithContext(ctx).Find(&result).Where("user_id = @user_id", sql.Named("user_id", user_id)).Error; err != nil {
		config.GetConnection().Rollback()
	}
	config.GetConnection().Commit()

	return result, nil
}

func SelectWishlistById(ctx context.Context, wishlist_id string) (Wishlist, error) {
	var result Wishlist

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).First(&result, "wishlist_id = @wishlist_id", sql.Named("wishlist_id", wishlist_id)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return result, nil
}

func UpdateWishlist(ctx context.Context, data Wishlist, wishlist_id string) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Model(data).Where("wishlist_id = @wishlist_id", sql.Named("wishlist_id", wishlist_id)).Updates(&Wishlist{WishlistId: wishlist_id, UserId: data.UserId, ProductId: data.ProductId, Quantity: data.Quantity}).Error; err != nil {
		config.GetConnection().Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func DeleteWishlistById(ctx context.Context, wishlist_id string) error {
	var result Wishlist

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Delete(&result, "wishlist_id = @wishlist_id", sql.Named("wishlist_id", wishlist_id)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func DeleteWishlistByStore(ctx context.Context, store_name string) error {
	var result Wishlist

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Delete(&result, "store_name = @store_name", sql.Named("store_name", store_name)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func DeleteWishlistByUser(ctx context.Context, user_id string) error {
	var result Wishlist

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Delete(&result, "user_id = @user_id", sql.Named("user_id", user_id)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}
