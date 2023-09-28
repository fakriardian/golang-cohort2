package dto

import (
	"fmt"
	"strconv"
	"strings"
)

type People []Person

func (people *People) Find(search string, key string) (Person, error) {
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
		}
	}

	return Person{}, fmt.Errorf("%s not found or out of range", search)
}
