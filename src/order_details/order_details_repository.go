package order_details

import (
	"config"
	"context"
	"database/sql"
)

func InsertOrderDetail(ctx context.Context, data interface{}) error {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Create(data).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return nil
}

func SelectOrderDetailById(ctx context.Context, order_id int) (OrderDetail, error) {
	var result OrderDetail

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).First(&result, "order_id = @order_id", sql.Named("order_id", order_id)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
	}
	config.GetConnection().WithContext(ctx).Commit()

	return result, nil
}
