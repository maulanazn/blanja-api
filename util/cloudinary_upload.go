package util

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/spf13/viper"
)

func UploadCloudinary(file multipart.File) (*uploader.UploadResult, error) {
  var viper *viper.Viper = LoadConfig("./../", "blanja.yaml", "yaml")
	cld, err := cloudinary.NewFromURL(viper.GetString("thirdparty.cloudinaryurl"))
	if err != nil {
		panic(err)
	}
	timeout, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	responseimage, err := cld.Upload.Upload(timeout, file, uploader.UploadParams{Folder: viper.GetString("thirdparty.cloudinaryfolder")})
	defer cancel()
	defer file.Close()
	if err != nil {
		fmt.Println("Time out")
	}

	return responseimage, nil
}
