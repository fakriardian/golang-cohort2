package main

import (
	"be/dto"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Users dto.Users
	PORT  string
)

func init() {
	data, err := os.ReadFile("./file.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
	}

	if err := json.Unmarshal(data, &Users); err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
}
