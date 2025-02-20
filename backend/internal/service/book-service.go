package service

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/aws-cakap-intern/book-store/internal/dto"
	"github.com/aws-cakap-intern/book-store/internal/entity"
	"github.com/aws-cakap-intern/book-store/internal/http/binder"
	"github.com/aws-cakap-intern/book-store/internal/repository"
	"github.com/aws-cakap-intern/book-store/pkg/execption"
	"github.com/google/uuid"
)

type BookService interface {
	GetBooks() ([]*dto.BookResponse, *execption.ApiExecption)
	GetBook(bookID string) (*dto.BookResponse, *execption.ApiExecption)
	CreateBook(input binder.CreateBook, categoryIDS []uint, file multipart.File, fileHeader *multipart.FileHeader) (*dto.BookResponse, *execption.ApiExecption)
	UpdateBook(input binder.UpdateBook, categoryIDS []uint, file multipart.File, fileHeader *multipart.FileHeader) (*dto.BookResponse, *execption.ApiExecption)
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
func (b *bookService) CreateBook(input binder.CreateBook, categoryIDS []uint, file multipart.File, fileHeader *multipart.FileHeader) (*dto.BookResponse, *execption.ApiExecption) {
	var categories []*entity.Category
	err := b.categoryRepo.FindByIDs(categoryIDS, &categories)
	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, "Error retrieving categories")
	}

	if len(categories) != len(categoryIDS) {
		return nil, execption.NewApiExecption(http.StatusBadRequest, "Some category IDs do not exist")
	}

	imagePath := ""
	if file != nil && fileHeader != nil {
		imagePath, err = b.saveFile(file, fileHeader)
		if err != nil {
			return nil, execption.NewApiExecption(http.StatusInternalServerError, "Error saving image")
		}
	}

	book := &entity.Book{
		Title:       input.Title,
		Price:       input.Price,
		ImagePath:   imagePath,
		Description: input.Description,
		Categories:  convertCategories(categories),
	}

	book, err = b.bookRepo.Create(book, categoryIDS)
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

	// Delete the image file if it exists
	if book.ImagePath != "" {
		imagePath := "uploads/" + book.ImagePath // Ensure correct path
		err := os.Remove(imagePath)
		if err != nil && !os.IsNotExist(err) {
			fmt.Println("Error deleting image file:", err)
		}
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
func (b *bookService) UpdateBook(input binder.UpdateBook, categoryIDS []uint, file multipart.File, fileHeader *multipart.FileHeader) (*dto.BookResponse, *execption.ApiExecption) {
	bookID, err := strconv.ParseUint(input.ID, 10, 0)
	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	var categories []*entity.Category
	err = b.categoryRepo.FindByIDs(categoryIDS, &categories)
	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, "Error retrieving categories")
	}

	// Ensure all requested categories exist
	if len(categories) != len(categoryIDS) {
		return nil, execption.NewApiExecption(http.StatusBadRequest, "Some category IDs do not exist")
	}

	book, err := b.bookRepo.GetById(uint(bookID))
	if err != nil {
		return nil, execption.NewApiExecption(http.StatusNotFound, "Book not found")
	}

	// Handle image update (delete old file if a new one is provided)
	if file != nil && fileHeader != nil {
		// Delete old image file if it exists
		if book.ImagePath != "" {
			_ = os.Remove(book.ImagePath) // Ignore errors in deletion
		}

		// Save new file
		imagePath, err := b.saveFile(file, fileHeader)
		if err != nil {
			return nil, execption.NewApiExecption(http.StatusInternalServerError, "Error saving image")
		}
		book.ImagePath = imagePath
	}

	updatedBook := &entity.Book{
		ID:          uint(bookID),
		Title:       input.Title,
		Price:       input.Price,
		ImagePath:   book.ImagePath,
		Description: input.Description,
		Categories:  convertCategories(categories), // Assign updated categories
	}

	book, err = b.bookRepo.Update(updatedBook, categoryIDS)
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

func convertCategories(categories []*entity.Category) []entity.Category {
	var result []entity.Category
	for _, category := range categories {
		result = append(result, *category)
	}
	return result
}

func (b *bookService) saveFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	dir := "uploads"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", err
	}

	// Generate a unique file name using timestamp and UUID
	uniqueName := fmt.Sprintf("%d_%s%s", time.Now().Unix(), uuid.New().String(), filepath.Ext(fileHeader.Filename))
	filePath := filepath.Join(dir, uniqueName)

	// Normalize to forward slashes for JSON compatibility
	filePath = strings.ReplaceAll(filePath, "\\", "/")

	// Create the new file
	outFile, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	// Copy file data
	_, err = outFile.ReadFrom(file)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
