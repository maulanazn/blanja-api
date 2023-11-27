package service

import (
	"context"
	"entity"
	"fmt"
	"log"
	"net/http"
	"repository"
	"response"
	"util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	productImage, productImageHeader, err := req.FormFile("image")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := util.ValidateImage(productImage, productImageHeader, writer); err != nil {
		failedResponse := response.ToWebResponse(400, err.Error())
		_, err := fmt.Fprint(writer, failedResponse)
		if err != nil {
			log.Println(err)
		}
		return
	}

	responseImage, responseImageErr := util.UploadCloudinary(productImage)
	util.BadStatusIfError(responseImageErr, writer)

	userid, userIdErr := req.Cookie("USR_ID")
	util.PanicIfError(userIdErr)

	products := &entity.Products{
		ProductId:   primitive.NewObjectID(),
		UserId:      userid.Value,
		Image:       responseImage.SecureURL,
		ProductName: req.FormValue("product_name"),
		Rating:      util.ConvertStrInt(req.FormValue("rating"), 10, 64),
		Price:       util.ConvertStrInt(req.FormValue("price"), 10, 64),
		Quantity:    util.ConvertStrInt(req.FormValue("quantity"), 10, 64),
	}

	if err := repository.CreateProduct(ctx, *products); err != nil {
		writer.WriteHeader(500)
		if _, err := writer.Write([]byte("Failed to insert, check again later")); err != nil {
			log.Println(err)
		}
		return
	}

	writer.WriteHeader(201)
	registerResponse := response.ToWebResponse(201, "Successfully create products")
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

	if err := util.ValidateImage(productImage, productImageHeader, writer); err != nil {
		failedResponse := response.ToWebResponse(400, err.Error())
		if _, err := fmt.Fprint(writer, failedResponse); err != nil {
			log.Println(err)
		}
		return
	}

	responseImage, responseImageErr := util.UploadCloudinary(productImage)
	util.BadStatusIfError(responseImageErr, writer)

	products := &entity.Products{
		Image:       responseImage.SecureURL,
		ProductName: req.FormValue("product_name"),
		Rating:      util.ConvertStrInt(req.FormValue("rating"), 10, 64),
		Price:       util.ConvertStrInt(req.FormValue("price"), 10, 64),
		Quantity:    util.ConvertStrInt(req.FormValue("quantity"), 10, 64),
	}

	if queryErr := repository.UpdateProduct(ctx, queryParam.Get("id"), *products); queryErr != nil {
 		writer.WriteHeader(400)
		if _, err := writer.Write([]byte("Failed to update, check again later")); err != nil {
			log.Println(err)
		}
	 	return
	}

	writer.WriteHeader(201)
	updateProductResponse := response.ToWebResponse(201, "Successfully update product")
	if _, err := fmt.Fprint(writer, updateProductResponse); err != nil {
		log.Println(err)
	}
}

func GetProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	queryParam := req.URL.Query()
  userid, userIdErr := req.Cookie("USR_ID")
  if userIdErr != nil {
    log.Println(userIdErr.Error())
  }

	if queryParam.Has("id") {
    result := repository.SelectProduct(ctx, queryParam.Get("id"))

    productIdResponse := response.ToGetProduct(200, "Success get requested product", response.GetProduct{
      Data: *result,
    })
		
		_, err := fmt.Fprint(writer, productIdResponse);
		if err != nil {
			log.Println(err.Error())
		}
		return
	}	
	  
  var products []entity.Products

	cursor := repository.SelectUserProduct(ctx, userid.Value)

  if err := cursor.All(ctx, &products); err != nil {
    log.Println(writer.Write([]byte("Failed to get product data")))
    return
  }

  userProductResponse := response.ToGetProducts(200, "Success get all product", response.GetProducts{
    Data: products,
  })
	if _, err := fmt.Fprint(writer, userProductResponse); err != nil {
		log.Println(err.Error())
	}
}
