package service

import (
	"context"
	"entity"
	"fmt"
	"helper"
	"net/http"
	"repository"
	"response"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type attribute struct {
	Name        string
	Description string
}

var validate = validator.New()

func AddProduct(ctx context.Context, productreq entity.Products, writer http.ResponseWriter, request *http.Request) {
	attr := attribute{}
	userid, err := request.Cookie("USR_ID")
	helper.PanicIfError(err)
	if validateerr := validate.Struct(&productreq); validateerr != nil {
		writer.WriteHeader(400)
		writer.Write([]byte(validateerr.Error()))
		return
	}

	products := entity.Products{
		ProductId: primitive.NewObjectID(),
		UserId:    userid.Value,
		CategoryName: entity.CategoryName{
			Name:        attr.Name,
			Description: attr.Description,
		},
		Image:       productreq.Image,
		ProductName: productreq.ProductName,
		Brand: entity.BrandName{
			Name:        attr.Name,
			Description: attr.Description,
		},
		Rating: productreq.Rating,
		Price:  productreq.Price,
		Color: entity.ColorName{
			Name:        attr.Name,
			Description: attr.Description,
		},
		Size: entity.SizeName{
			Name:        attr.Name,
			Description: attr.Description,
		},
		Quantity: productreq.Quantity,
	}

	if err := repository.CreateProduct(context.TODO(), products); err != nil {
		writer.WriteHeader(500)
		writer.Write([]byte("Failed to insert, check again later"))
		return
	}

	writer.WriteHeader(201)
	registerResponse := response.ToGetProducts(201, "Successfully create addresss", products)
	fmt.Fprint(writer, registerResponse)
}
