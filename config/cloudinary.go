package config

import (
	cloudinary "github.com/cloudinary/cloudinary-go/v2"
)

func GetCloudinaryConfig() *cloudinary.Cloudinary {
	config := GetConfig()
	cld, err := cloudinary.NewFromURL(config.GetString("CLOUDINARY_URL"))
	if err != nil {
		panic(err)
	}

	return cld
}
