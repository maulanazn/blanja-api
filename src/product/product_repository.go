package product

import (
	"config"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"util"
)

func InsertProduct(ctx context.Context, product interface{}, writer http.ResponseWriter) {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Create(product).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
		writer.WriteHeader(400)
		if _, err := writer.Write([]byte("Failed to insert, check again later")); err != nil {
			log.Println(err)
			return
		}
		return
	}
	config.GetConnection().WithContext(ctx).Commit()
}

func UpdateProduct(ctx context.Context, product Products, product_id string, writer http.ResponseWriter) {
	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Model(product).Where("product_id = @product_id", sql.Named("product_id", product_id)).Updates(&Products{
		Image:       product.Image,
		ProductName: product.ProductName,
		StoreName:   product.StoreName,
		Rating:      product.Rating,
		Price:       product.Price,
		Quantity:    product.Quantity,
	}).Error; err != nil {
		config.GetConnection().Rollback()
		writer.WriteHeader(400)
		if _, err := writer.Write([]byte("Failed to update, check again later")); err != nil {
			log.Println(err)
		}

		return
	}
	config.GetConnection().WithContext(ctx).Commit()
}

/*
* @description Get Product By Store and User
* @params context, string
* @return slice of product (pointer)
 */
func SelectProductByStore(ctx context.Context, store_name string) (*[]Products, error) {
	var results *[]Products

	config.GetConnection().Begin()
	if err := config.GetConnection().WithContext(ctx).Find(&results).Where("store_name = @store_name", sql.Named("store_name", store_name)).Error; err != nil {
		config.GetConnection().Rollback()
	}
	config.GetConnection().Commit()

	return results, nil
}

func SelectProductByUser(ctx context.Context, user_id string) (*[]Products, error) {
	var results *[]Products

	config.GetConnection().Begin()
	if err := config.GetConnection().WithContext(ctx).Find(&results).Where("user_id = @user_id", sql.Named("user_id", user_id)).Error; err != nil {
		config.GetConnection().Rollback()
	}
	config.GetConnection().Commit()

	return results, nil
}

/*
* @description Get Product By Id
* @params context, string
* @return Single product (pointer)
 */
func SelectProductById(ctx context.Context, product_id string) (*Products, error) {
	var result *Products

	config.GetConnection().Begin()
	if err := config.GetConnection().WithContext(ctx).First(&result, "product_id = @product_id", sql.Named("product_id", product_id)).Error; err != nil {
		config.GetConnection().Rollback()
	}
	config.GetConnection().Commit()

	return result, nil
}

/*
* @description Delete Product By Id, Store, and User
* @params context, string, http.ResponseWriter
* @return None
 */
func DeleteProductById(ctx context.Context, product_id string, writer http.ResponseWriter) {
	var result *Products

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Delete(&result, "product_id = @product_id", sql.Named("product_id", product_id)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
		_, err := fmt.Fprint(writer, util.ToWebResponse(400, err.Error()))
		if err != nil {
			util.Log2File(err.Error())
		}

		return
	}

	config.GetConnection().WithContext(ctx).Commit()
}

func DeleteProductByStore(ctx context.Context, store_name string, writer http.ResponseWriter) {
	var result *Products

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Delete(&result, "store_name = @store_name", sql.Named("store_name", store_name)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
		_, err := fmt.Fprint(writer, util.ToWebResponse(400, err.Error()))
		if err != nil {
			util.Log2File(err.Error())
		}

		return
	}
	config.GetConnection().WithContext(ctx).Commit()
}

func DeleteProductByUser(ctx context.Context, user_id string, writer http.ResponseWriter) {
	var result *Products

	config.GetConnection().WithContext(ctx).Begin()
	if err := config.GetConnection().WithContext(ctx).Delete(&result, "user_id = @user_id", sql.Named("user_id", user_id)).Error; err != nil {
		config.GetConnection().WithContext(ctx).Rollback()
		_, err := fmt.Fprint(writer, util.ToWebResponse(400, err.Error()))
		if err != nil {
			util.Log2File(err.Error())
		}

		return
	}
	config.GetConnection().WithContext(ctx).Commit()
}
