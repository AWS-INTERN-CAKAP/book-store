package binder

type GetCategory struct {
	ID string `param:"id" validate:"required"`
}

type CreateCategory struct {
	Name string `json:"name" validate:"required"`
}

type UpdateCategory struct {
	ID   string `param:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type DeleteCategory struct {
	ID string `param:"id" validate:"required"`
}