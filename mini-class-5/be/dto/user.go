package dto

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	ID        int    `json:"ID"`
	Nama      string `json:"name"`
	Email     string `json:"email"`
	Alamat    string `json:"alamat"`
	Pekerjaan string `json:"pekerjaan"`
	Alasan    string `json:"alasan"`
}

func (user User) Print() {
	refVal := reflect.ValueOf(user)
	typeOf := refVal.Type()

	for i := 0; i < refVal.NumField(); i++ {
		key := typeOf.Field(i).Name
		value := refVal.Field(i).Interface()
		fmt.Printf("%s: %v\n", key, value)
	}
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Users []User

func (people *Users) Find(search string, key string) (User, error) {
	for _, data := range *people {
		if key == "nama" {
			if strings.EqualFold(data.Nama, search) {
				return data, nil
			}
		} else if key == "ID" {
			searchConv, _ := strconv.Atoi(search)
			if data.ID == searchConv {
				return data, nil
			}
		} else if key == "email" {
			if data.Email == search {
				return data, nil
			}
		}
	}

	return User{}, fmt.Errorf("%s not found or out of range", search)
}
