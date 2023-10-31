package helper

import (
	"belanjabackend/config"
	"bytes"
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadCloudinary(file string) (*uploader.UploadResult, error) {
	osfile, osfileerr := os.Open(file)
	PanicIfError(osfileerr)

	bytesNew := bytes.NewBuffer(nil)
	io.Copy(bytesNew, osfile)

	responseimage, err := config.GetCloudinaryConfig().Upload.Upload(context.Background(), bytesNew, uploader.UploadParams{Folder: "belanja"})
	PanicIfError(err)

	return responseimage, nil
}
