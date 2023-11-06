package main

import (
	"final-challenge/config"
	"final-challenge/routes"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if isContainer, _ := strconv.ParseBool(os.Getenv("USING_CONTAINTER")); !isContainer {
		if err := godotenv.Load(".env"); err != nil {
			log.Println("Error loading .env file")
		}
	}

	deps := config.Init()

	router := gin.Default()
	router.SetTrustedProxies([]string{"*"})

	gin.SetMode(os.Getenv("ENV"))

	routes.InitAuthRoutes(deps.DB, router)
	routes.InitProductRoutes(deps, router)

	log.Fatal(router.Run(":" + os.Getenv("PORT")))
}
