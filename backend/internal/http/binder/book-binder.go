package binder

import "mime/multipart"

type GetBook struct {
	ID string `param:"id" validate:"required"`
}

type CreateBook struct {
	Title       string                `form:"title" binding:"required"`
	Price       int               `form:"price" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Categories  []uint                `form:"categories" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
}

type UpdateBook struct {
	ID          string                `param:"id" validate:"required"`
	Title       string                `form:"title" binding:"required"`
	Price       int               `form:"price" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Categories  []uint                `form:"categories" binding:"required"`
	Image       *multipart.FileHeader `form:"image" binding:"required"`
}

type DeleteBook struct {
	ID string `param:"id" validate:"required"`
}
