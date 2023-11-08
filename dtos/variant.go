package dtos

import (
	"time"
)

type Variant struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name"`
	Quantity  int        `json:"quantity"`
	ProductID string     `json:"productId"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type VariantCreateRequest struct {
	Name      string `json:"name" form:"name" binding:"required"`
	Quantity  int    `json:"quantity" form:"quantity" binding:"required"`
	ProductID string `json:"productId" form:"productId" binding:"required,uuid"`
	AdminID   string `binding:"required"`
}

type VariantUpdateRequest struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Quantity int    `json:"quantity" form:"quantity" binding:"-"`
	AdminID  string `binding:"required"`
}

type VariantIDUri struct {
	ID string `uri:"id" binding:"required,uuid"`
}
