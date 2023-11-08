package dtos

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	Name      string     `json:"name"`
	ImageUrl  string     `json:"imageUrl"`
	UserID    uuid.UUID  `json:"-"`
	User      User       `json:"admin"`
	Variants  []Variant  `json:"variants"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type ProductCreateRequest struct {
	Name    string                `form:"name" binding:"required"`
	File    *multipart.FileHeader `form:"file" binding:"required"`
	AdminID string                `binding:"required"`
}

type ProductUpdateRequest struct {
	Name    string                `form:"name" binding:"required"`
	File    *multipart.FileHeader `form:"file" binding:"-"`
	AdminID string                `binding:"required"`
}

type ProductIDUri struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type CreateProductResponse struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name"`
	ImageUrl  string     `json:"imageUrl"`
	UserID    string     `json:"adminId"`
	Variants  []Variant  `json:"variants"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
