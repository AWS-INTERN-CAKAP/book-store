package service

import (
	"net/http"
	"strconv"

	"github.com/aws-cakap-intern/book-store/internal/dto"
	"github.com/aws-cakap-intern/book-store/internal/entity"
	"github.com/aws-cakap-intern/book-store/internal/http/binder"
	"github.com/aws-cakap-intern/book-store/internal/repository"
	"github.com/aws-cakap-intern/book-store/pkg/execption"
)

type BookService interface {
	GetBooks() ([]*dto.BookResponse, *execption.ApiExecption)
	GetBook(bookID string) (*dto.BookResponse, *execption.ApiExecption)
	CreateBook(input binder.CreateBook) (*dto.BookResponse, *execption.ApiExecption)
	UpdateBook(input binder.UpdateBook) (*dto.BookResponse, *execption.ApiExecption)
	DeleteBook(bookID string) *execption.ApiExecption
}

type bookService struct {
	bookRepo     repository.BookRepository
	categoryRepo repository.CategoryRepository
}

func NewBookService(bookRepo repository.BookRepository, categoryRepo repository.CategoryRepository) BookService {
	return &bookService{bookRepo: bookRepo, categoryRepo: categoryRepo}
}

// CreateBook implements BookService.
func (b *bookService) CreateBook(input binder.CreateBook) (*dto.BookResponse, *execption.ApiExecption) {

	var categories []*entity.Category
	err := b.categoryRepo.FindByIDs(input.Categories, &categories)
	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, "Error retrieving categories")
	}

	// Ensure all requested categories exist
	if len(categories) != len(input.Categories) {
		return nil, execption.NewApiExecption(http.StatusBadRequest, "Some category IDs do not exist")
	}

	// Convert []*entity.Category to []entity.Category
	var categoryValues []entity.Category
	for _, category := range categories {
		categoryValues = append(categoryValues, *category)
	}

	book := &entity.Book{
		Title:       input.Title,
		Price:       input.Price,
		ImagePath:   input.ImagePath,
		Description: input.Description,
		Categories:  categoryValues,
	}

	book, err = b.bookRepo.Create(book, input.Categories)

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	response := &dto.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Price:       book.Price,
		ImagePath:   book.ImagePath,
		Description: book.Description,
		Categories:  []dto.CategoryResponse{},
		CreatedAt:   book.CreatedAt.String(),
		UpdatedAt:   book.UpdatedAt.String(),
	}

	for _, category := range categories {
		categoryResponse := dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		}
		response.Categories = append(response.Categories, categoryResponse)
	}

	return response, nil
}

// DeleteBook implements BookService.
func (b *bookService) DeleteBook(bookID string) *execption.ApiExecption {
	uintID, err := strconv.ParseUint(bookID, 10, 0)
	if err != nil {
		return execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	// Check if the book exists before deletion
	book, err := b.bookRepo.GetById(uint(uintID))
	if err != nil {
		if err == repository.ErrBookNotFound {
			return execption.NewApiExecption(http.StatusNotFound, "Book not found")
		}
		return execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	// Remove category associations (if applicable)
	if len(book.Categories) > 0 {
		err = b.bookRepo.Delete(uint(uintID)) // Ensure this function exists in the repository
		if err != nil {
			return execption.NewApiExecption(http.StatusInternalServerError, "Failed to remove category associations")
		}
	}

	// Proceed with deletion
	err = b.bookRepo.Delete(uint(uintID))
	if err != nil {
		return execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	return nil
}

// GetBook implements BookService.
func (b *bookService) GetBook(bookID string) (*dto.BookResponse, *execption.ApiExecption) {
	uintID, err := strconv.ParseUint(bookID, 10, 0)

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	book, err := b.bookRepo.GetById(uint(uintID))

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusNotFound, err.Error())
	}

	// Convert categories to response format
	var categoryResponses []dto.CategoryResponse
	for _, category := range book.Categories {
		categoryResponses = append(categoryResponses, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	response := &dto.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Price:       book.Price,
		ImagePath:   book.ImagePath,
		Description: book.Description,
		Categories:  categoryResponses,
		CreatedAt:   book.CreatedAt.String(),
		UpdatedAt:   book.UpdatedAt.String(),
	}

	return response, nil
}

// GetBooks implements BookService.
func (b *bookService) GetBooks() ([]*dto.BookResponse, *execption.ApiExecption) {
	books, err := b.bookRepo.GetAll()
	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	if len(books) == 0 {
		return []*dto.BookResponse{}, nil
	}

	var responses []*dto.BookResponse

	for _, book := range books {
		// Convert categories to response format
		var categoryResponses []dto.CategoryResponse
		for _, category := range book.Categories {
			categoryResponses = append(categoryResponses, dto.CategoryResponse{
				ID:   category.ID,
				Name: category.Name,
			})
		}

		responses = append(responses, &dto.BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			Price:       book.Price,
			ImagePath:   book.ImagePath,
			Description: book.Description,
			CreatedAt:   book.CreatedAt.String(),
			UpdatedAt:   book.UpdatedAt.String(),
			Categories:  categoryResponses, // Include categories in response
		})
	}

	return responses, nil
}

// UpdateBook implements BookService.
func (b *bookService) UpdateBook(input binder.UpdateBook) (*dto.BookResponse, *execption.ApiExecption) {
	bookID, err := strconv.ParseUint(input.ID, 10, 0)
	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	// Retrieve categories by their IDs
	var categories []*entity.Category
	err = b.categoryRepo.FindByIDs(input.Categories, &categories)
	if err != nil {
		return nil, execption.NewApiExecption(http.StatusBadRequest, "Invalid category IDs")
	}

	// Convert []*entity.Category to []entity.Category
	var categoryValues []entity.Category
	for _, category := range categories {
		categoryValues = append(categoryValues, *category)
	}

	book := &entity.Book{
		ID:          uint(bookID),
		Title:       input.Title,
		Price:       input.Price,
		ImagePath:   input.ImagePath,
		Description: input.Description,
		Categories:  categoryValues, // Assign updated categories
	}

	book, err = b.bookRepo.Update(book, input.Categories)
	if err != nil {
		if err == repository.ErrBookNotFound {
			return nil, execption.NewApiExecption(http.StatusNotFound, err.Error())
		}
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	// Convert categories to response format
	var categoryResponses []dto.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, dto.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	response := &dto.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Price:       book.Price,
		ImagePath:   book.ImagePath,
		Description: book.Description,
		CreatedAt:   book.CreatedAt.String(),
		UpdatedAt:   book.UpdatedAt.String(),
		Categories:  categoryResponses, // Include categories in response
	}

	return response, nil
}
