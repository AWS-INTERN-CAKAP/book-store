package binder

import "mime/multipart"

type GetBook struct {
	ID string `param:"id" validate:"required"`
}

type CreateBook struct {
	Title       string                `form:"title" validate:"required"`
	Price       int                   `form:"price" validate:"required"`
	Description string                `form:"description" validate:"required"`
	Image       *multipart.FileHeader `form:"image" validate:"required"`
}

type UpdateBook struct {
	ID          string                `param:"id" validate:"required"`
	Title       string                `form:"title" validate:"required"`
	Price       int                   `form:"price" validate:"required"`
	Description string                `form:"description" validate:"required"`
	Image       *multipart.FileHeader `form:"image"`
}

type DeleteBook struct {
	ID string `param:"id" validate:"required"`
}
