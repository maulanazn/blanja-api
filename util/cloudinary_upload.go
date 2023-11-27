package util

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadCloudinary(file multipart.File) (*uploader.UploadResult, error) {
  // var viper *viper.Viper = LoadConfig("./../", "blanja.yaml", "yaml")
	// cld, err := cloudinary.NewFromURL(viper.GetString("thirdparty.cloudinaryurl"))
	/* if err != nil {
	*	panic(err)
	}*/
	cld, err := cloudinary.NewFromURL("cloudinary://318565578158456:FKmFXz621H8AlXqgfaf9p55b-aQ@dra9sesmi")
	if err != nil {
		panic(err)
	}
	timeout, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	responseimage, err := cld.Upload.Upload(timeout, file, uploader.UploadParams{Folder: "belanja"})
	defer cancel()
	defer file.Close()
	if err != nil {
		fmt.Println("Time out")
	}

	return responseimage, nil
}
