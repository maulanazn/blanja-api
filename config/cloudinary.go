package config

import (
	cloudinary "github.com/cloudinary/cloudinary-go/v2"
)

func GetCloudinaryConfig() *cloudinary.Cloudinary {
	cld, err := cloudinary.NewFromURL(GetCLDURL())
	if err != nil {
		panic(err)
	}

	return cld
}
