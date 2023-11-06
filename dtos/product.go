package dtos

import "mime/multipart"

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
