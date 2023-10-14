package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Tolong masukan nama atau nomer absesn!")
		fmt.Println("Contoh: 'go run . Fitri' atau 'go run . 2'")
		return
	}

	argument := os.Args[1]
	tipe := "nama"

	if _, err := strconv.Atoi(argument); err == nil {
		tipe = "ID"
	}

	response, err := People.Find(argument, tipe)
	if err != nil {
		fmt.Println(err)
		return
	}

	response.Print()
}
