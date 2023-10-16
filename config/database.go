package config

import (
	"log"
	"mini-challenge/models"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	mysqlInfo := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Order{}, &models.Item{})
	if err != nil {
		log.Fatalln("Error occurred during auto migration:", err)
	} else {
		log.Println("Successfully migrate database")
	}

	log.Println("Successfully connected to database")
	return db, nil
}
