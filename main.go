package main

import (
	"log"
	"mini-challenge/config"
	"mini-challenge/routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}

	db, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.SetTrustedProxies([]string{"*"})

	env := os.Getenv("ENV")
	gin.SetMode(env)

	routes.InitOrderRoutes(db, router)

	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
