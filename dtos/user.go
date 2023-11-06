package dtos

type UserRegister struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email" form:"email"`
	Password string `json:"password" binding:"required,gte=6" form:"password"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email" form:"email"`
	Password string `json:"password" binding:"required" form:"password"`
}
