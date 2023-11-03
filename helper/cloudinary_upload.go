package helper

import (
	"config"
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadCloudinary(file multipart.File) (*uploader.UploadResult, error) {
	timeout, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	responseimage, err := config.GetCloudinaryConfig().Upload.Upload(timeout, file, uploader.UploadParams{Folder: os.Getenv("CLOUDINARY_FOLDER")})
	defer cancel()
	defer file.Close()
	if err != nil {
		fmt.Println("Time out")
	}

	return responseimage, nil
}
