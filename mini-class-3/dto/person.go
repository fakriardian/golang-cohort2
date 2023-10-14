package dto

import (
	"fmt"
	"reflect"
)

type Person struct {
	ID        int    `json:"ID"`
	Nama      string `json:"name"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Alasan    string `json:"alasan"`
}

func (person Person) Print() {
	refVal := reflect.ValueOf(person)
	typeOf := refVal.Type()

	for i := 0; i < refVal.NumField(); i++ {
		key := typeOf.Field(i).Name
		value := refVal.Field(i).Interface()
		fmt.Printf("%s: %v\n", key, value)
	}
}
