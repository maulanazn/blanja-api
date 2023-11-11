package config

import (
	cloudinary "github.com/cloudinary/cloudinary-go/v2"
)

func GetCloudinaryConfig() *cloudinary.Cloudinary {
	// cld, err := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	cld, err := cloudinary.NewFromURL("cloudinary://318565578158456:FKmFXz621H8AlXqgfaf9p55b-aQ@dra9sesmi")
	if err != nil {
		panic(err)
	}

	return cld
}
