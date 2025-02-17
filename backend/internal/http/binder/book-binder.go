package binder

type GetBook struct {
	ID string `param:"id" validate:"required"`
}

type CreateBook struct {
	Title string `json:"name" validate:"required"`
	Price string `json:"price" validate:"required"`
	ImagePath string `json:"imagePath" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateBook struct {
	ID   string `param:"id" validate:"required"`
	Title string `json:"name" validate:"required"`
	Price string `json:"price" validate:"required"`
	ImagePath string `json:"imagePath" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type DeleteBook struct {
	ID string `param:"id" validate:"required"`
}