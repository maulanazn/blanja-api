package product

import (
	"context"
)

func InsertProduct(ctx context.Context, product Products) error {
	return nil
}

func UpdateProduct(ctx context.Context, id string, product Products) error {
	return nil
}

func SelectProductByUser(ctx context.Context, userId string) (*[]Products, error) {
	var results *[]Products

	return results, nil
}

func SelectProductById(ctx context.Context, id string) (*Products, error) {
	var result *Products

	return result, nil
}
