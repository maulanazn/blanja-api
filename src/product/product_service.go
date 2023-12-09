package product

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"util"
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
		UserId:      userid,
		Image:       responseImage.SecureURL,
		ProductName: req.FormValue("product_name"),
		StoreName:   req.FormValue("store_name"),
		Rating:      util.ConvertStrInt(req.FormValue("rating"), 10, 64),
		Price:       util.ConvertStrInt(req.FormValue("price"), 10, 64),
		Quantity:    util.ConvertStrInt(req.FormValue("quantity"), 10, 64),
	}

	InsertProduct(ctx, products, writer)

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

	UpdateProduct(ctx, *products, queryParam.Get("id"), writer)

	writer.WriteHeader(201)
	if _, err := fmt.Fprint(writer, util.ToWebResponse(201, "Successfully update product")); err != nil {
		log.Println(err)
	}
}

func GetProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	queryParam := req.URL.Query()
	userid := util.DecodeToken(req.Header.Get("Authorization"))

	if queryParam.Has("product_id") {
		result, _ := SelectProductById(ctx, queryParam.Get("product_id"))

		_, err := fmt.Fprint(writer, ToGetProduct(200, "Success get requested product", GetProductStruct{
			Data: *result,
		}))
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	if queryParam.Has("store_name") {
		result, _ := SelectProductByStore(ctx, queryParam.Get("store_name"))

		_, err := fmt.Fprint(writer, ToGetProducts(200, "Success get requested product", GetProducts{
			Data: *result,
		}))
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	results, err := SelectProductByUser(ctx, userid)
	if err != nil {
		log.Println(err.Error())
	}

	if _, err := fmt.Fprint(writer, ToGetProducts(200, "Success get all product", GetProducts{
		Data: *results,
	})); err != nil {
		log.Println(err.Error())
	}
}

func DeleteProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	queryParam := req.URL.Query()
	userid := util.DecodeToken(req.Header.Get("Authorization"))

	if queryParam.Has("product_id") {
		DeleteProductById(ctx, queryParam.Get("product_id"), writer)

		_, err := fmt.Fprint(writer, util.ToWebResponse(202, "Success deleted the requested product"))
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	if queryParam.Has("store_name") {
		DeleteProductByStore(ctx, queryParam.Get("store_name"), writer)

		_, err := fmt.Fprint(writer, util.ToWebResponse(202, "Success deleted the requested product"))
		if err != nil {
			log.Println(err.Error())
		}
		return
	}

	DeleteProductByUser(ctx, userid, writer)

	if _, err := fmt.Fprint(writer, util.ToWebResponse(202, "Success deleted the requested product")); err != nil {
		log.Println(err.Error())
	}
}
