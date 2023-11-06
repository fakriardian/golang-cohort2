package libs

import (
	"final-challenge/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDatabase struct {
	Host         string
	User         string
	Password     string
	Port         string
	DatabaseName string
}

func InitDatabase(opt *ConfigDatabase) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", opt.User, opt.Password, opt.Host, opt.Port, opt.DatabaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}

	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Variant{})
	if err != nil {
		log.Fatalln("Error occurred during auto migration:", err)
	} else {
		log.Println("Successfully migrate database")
	}

	log.Println("Successfully connected to database")
	return db
}
