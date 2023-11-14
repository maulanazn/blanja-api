package service

import (
	"context"
	"entity"
	"fmt"
	"net/http"
	"repository"
	"request"
	"response"
	"util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProduct(ctx context.Context, writer http.ResponseWriter, req *http.Request) {
	productreq := request.AddProductRequest{}
	if err := util.DecodeRequestAndValidate(writer, req, &productreq); err != nil {
		util.PanicIfError(err)
		return
	}

	userid, err := req.Cookie("USR_ID")
	util.PanicIfError(err)

	products := &entity.Products{
		ProductId:   primitive.NewObjectID(),
		UserId:      userid.Value,
		Image:       productreq.Image,
		ProductName: productreq.ProductName,
		Rating:      productreq.Rating,
		Price:       productreq.Price,
		Quantity:    productreq.Quantity,
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
