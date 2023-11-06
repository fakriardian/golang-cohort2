package config

import (
	"final-challenge/libs"
	"os"

	"gorm.io/gorm"
)

type Deps struct {
	DB   *gorm.DB
	STRG *libs.CloudinaryService
}

func Init() *Deps {
	return &Deps{
		DB: libs.InitDatabase(&libs.ConfigDatabase{
			Host:         os.Getenv("DB_HOST"),
			User:         os.Getenv("DB_USER"),
			Password:     os.Getenv("DB_PASSWORD"),
			Port:         os.Getenv("DB_PORT"),
			DatabaseName: os.Getenv("DB_NAME"),
		}),
		STRG: libs.InitCloudinary(&libs.ConfigCloudinary{
			CloudName:      os.Getenv("CLOUDINARY_CLOUD_NAME"),
			CloudApiKey:    os.Getenv("CLOUDINARY_API_KEY"),
			CloudApiSecret: os.Getenv("CLOUDINARY_API_SECRET"),
			CloudFolder:    os.Getenv("CLOUDINARY_UPLOAD_FOLDER"),
		}),
	}

}
