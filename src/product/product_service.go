package product

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	productImage, productImageHeader, err := req.FormFile("image")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	util.ValidateImage(productImage, productImageHeader, writer)

	responseImage, responseImageErr := util.UploadCloudinary(productImage)
	util.BadStatusIfError(responseImageErr, writer)

	userid := util.DecodeToken(req.Header.Get("Authorization"))

	products := &Products{
		ProductId:   primitive.NewObjectID(),
		UserId:      userid,
		Image:       responseImage.SecureURL,
		ProductName: req.FormValue("product_name"),
		StoreName:   req.FormValue("store_name"),
		Rating:      util.ConvertStrInt(req.FormValue("rating"), 10, 64),
		Price:       util.ConvertStrInt(req.FormValue("price"), 10, 64),
		Quantity:    util.ConvertStrInt(req.FormValue("quantity"), 10, 64),
	}

	if err := InsertProduct(ctx, *products); err != nil {
		writer.WriteHeader(500)
		if _, err := writer.Write([]byte("Failed to insert, check again later")); err != nil {
			log.Println(err)
		}
		return
	}

	writer.WriteHeader(201)
	registerResponse := util.ToWebResponse(201, "Successfully create products")
	if _, err := fmt.Fprint(writer, registerResponse); err != nil {
		log.Println(err)
	}
}

func EditProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	queryParam := req.URL.Query()
	productImage, productImageHeader, err := req.FormFile("image")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	util.ValidateImage(productImage, productImageHeader, writer)

	responseImage, responseImageErr := util.UploadCloudinary(productImage)
	util.BadStatusIfError(responseImageErr, writer)

	products := &Products{
		Image:       responseImage.SecureURL,
		ProductName: req.FormValue("product_name"),
		StoreName:   req.FormValue("store_name"),
		Rating:      util.ConvertStrInt(req.FormValue("rating"), 10, 64),
		Price:       util.ConvertStrInt(req.FormValue("price"), 10, 64),
		Quantity:    util.ConvertStrInt(req.FormValue("quantity"), 10, 64),
	}

	if queryErr := UpdateProduct(ctx, queryParam.Get("id"), *products); queryErr != nil {
		writer.WriteHeader(400)
		if _, err := writer.Write([]byte("Failed to update, check again later")); err != nil {
			log.Println(err)
		}
		return
	}

	writer.WriteHeader(201)
	updateProductResponse := util.ToWebResponse(201, "Successfully update product")
	if _, err := fmt.Fprint(writer, updateProductResponse); err != nil {
		log.Println(err)
	}
}

func GetProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	queryParam := req.URL.Query()
	userid := util.DecodeToken(req.Header.Get("Authorization"))

	if queryParam.Has("id") {
		result, _ := SelectProductById(ctx, queryParam.Get("id"))

		productIdResponse := ToGetProduct(200, "Success get requested product", GetProductStruct{
			Data: *result,
		})

		_, err := fmt.Fprint(writer, productIdResponse)
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	var products []Products

	_, err := SelectProductByUser(ctx, userid)
	if err != nil {
		log.Println(err.Error())
	}

	userProductResponse := ToGetProducts(200, "Success get all product", GetProducts{
		Data: products,
	})

	if _, err := fmt.Fprint(writer, userProductResponse); err != nil {
		log.Println(err.Error())
	}
}
