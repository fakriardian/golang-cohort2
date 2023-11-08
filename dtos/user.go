package dtos

import (
	"time"
)

type User struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	Products  []Product  `json:"-"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type UserRegister struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email" form:"email"`
	Password string `json:"password" binding:"required,gte=6" form:"password"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email" form:"email"`
	Password string `json:"password" binding:"required" form:"password"`
}
