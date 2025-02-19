package dto

type BookResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Price       int    `json:"price"`
	ImagePath   string `json:"imagePath"`
	Description string `json:"description"`
	Categories  []CategoryResponse `json:"categories"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}