package service

import (
	"context"
	"entity"
	"fmt"
	"net/http"
	"repository"
	"response"
	"util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	productimage, productimageheader, err := req.FormFile("image")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := util.ValidateImage(productimage, productimageheader, writer); err != nil {
		failedResponse := response.ToWebResponse(400, err.Error())
		fmt.Fprint(writer, failedResponse)
		return
	}

	responseimage, err := util.UploadCloudinary(productimage)
	util.BadStatusIfError(err, writer)

	userid, err := req.Cookie("USR_ID")
	util.PanicIfError(err)

	rating, ratingerr := util.ConvertStrInt(req.FormValue("rating"), 10, 64)
	util.PanicIfError(ratingerr)
	price, priceerr := util.ConvertStrInt(req.FormValue("price"), 10, 64)
	util.PanicIfError(priceerr)
	quantity, quantityerr := util.ConvertStrInt(req.FormValue("quantity"), 10, 64)
	util.PanicIfError(quantityerr)

	products := &entity.Products{
		ProductId:   primitive.NewObjectID(),
		UserId:      userid.Value,
		Image:       responseimage.SecureURL,
		ProductName: req.FormValue("product_name"),
		Rating:      rating,
		Price:       price,
		Quantity:    quantity,
	}

	if err := repository.CreateProduct(context.TODO(), *products); err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte("Failed to insert, check again later"))
		return
	}

	writer.WriteHeader(201)
	registerResponse := response.ToGetProducts(201, "Successfully create addresss", *products)
	fmt.Fprint(writer, registerResponse)
}

func EditProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	// queryParam := req.URL.Query()
	productimage, productimageheader, err := req.FormFile("image")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := util.ValidateImage(productimage, productimageheader, writer); err != nil {
		failedResponse := response.ToWebResponse(400, err.Error())
		fmt.Fprint(writer, failedResponse)
		return
	}

	responseimage, err := util.UploadCloudinary(productimage)
	util.BadStatusIfError(err, writer)

	userid, err := req.Cookie("USR_ID")
	util.PanicIfError(err)

	rating, ratingerr := util.ConvertStrInt(req.FormValue("rating"), 10, 64)
	util.PanicIfError(ratingerr)
	price, priceerr := util.ConvertStrInt(req.FormValue("price"), 10, 64)
	util.PanicIfError(priceerr)
	quantity, quantityerr := util.ConvertStrInt(req.FormValue("quantity"), 10, 64)
	util.PanicIfError(quantityerr)

	products := &entity.Products{
		ProductId:   primitive.NewObjectID(),
		UserId:      userid.Value,
		Image:       responseimage.SecureURL,
		ProductName: req.FormValue("product_name"),
		Rating:      rating,
		Price:       price,
		Quantity:    quantity,
	}

	// if err := repository.UpdateProduct(context.TODO(), byte(queryParam.Get("id").(byte)), *products); err != nil {
	// 	writer.WriteHeader(400)
	// 	writer.Write([]byte("Failed to update, check again later"))
	// 	return
	// }

	writer.WriteHeader(201)
	registerResponse := response.ToGetProducts(201, "Successfully create addresss", *products)
	fmt.Fprint(writer, registerResponse)
}

func GetProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {

}
