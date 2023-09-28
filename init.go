package main

import (
	"encoding/json"
	"fmt"
	"mini-class-3/dto"
	"os"
)

var (
	People dto.People
)

func init() {
	data, err := os.ReadFile("./file.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
	}

	if err := json.Unmarshal(data, &People); err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}
}
