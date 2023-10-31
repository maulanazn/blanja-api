package config

import (
	cloudinary "github.com/cloudinary/cloudinary-go/v2"
)

func GetCloudinaryConfig() *cloudinary.Cloudinary {
	cld, _ := cloudinary.NewFromParams("dra9sesmi", "318565578158456", "FKmFXz621H8AlXqgfaf9p55b-aQ")

	return cld
}
