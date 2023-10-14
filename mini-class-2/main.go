package main

import (
	"fmt"
)

func main() {
	kalimat := "selamat malam"

	count := make(map[string]int)

	for _, value := range kalimat {
		count[string(value)]++
		fmt.Printf("%c \n", value)
	}

	fmt.Printf("%v", count)
}
