package binder

type GetBook struct {
	ID string `param:"id" validate:"required"`
}

type CreateBook struct {
	Title       string  `json:"title" validate:"required"`
	Price       int     `json:"price" validate:"required"`
	ImagePath   string  `json:"imagePath" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Categories  []uint  `json:"categories" validate:"required"` // Accept category IDs
}

type UpdateBook struct {
	ID          string  `param:"id" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Price       int     `json:"price" validate:"required"`
	ImagePath   string  `json:"imagePath" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Categories  []uint  `json:"categories" validate:"required"` // Accept category IDs
}

type DeleteBook struct {
	ID string `param:"id" validate:"required"`
}
