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
	bookRepo repository.BookRepository
}

func NewBookService(bookRepo repository.BookRepository) BookService {
	return &bookService{bookRepo: bookRepo}
}

// CreateBook implements BookService.
func (b *bookService) CreateBook(input binder.CreateBook) (*dto.BookResponse, *execption.ApiExecption) {

	price, err := strconv.Atoi(input.Price)

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusBadRequest, "Invalid price format")
	}

	book := &entity.Book{
		Title:     input.Title,
		Price:     price,
		ImagePath: input.ImagePath,
		Description: input.Description,
	}

	book, err = b.bookRepo.Create(book)

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	response := &dto.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Price:       book.Price,
		ImagePath:   book.ImagePath,
		Description: book.Description,
		CreatedAt:   book.CreatedAt.String(),
		UpdatedAt:   book.UpdatedAt.String(),
	}

	return response, nil
}

// DeleteBook implements BookService.
func (b *bookService) DeleteBook(bookID string) *execption.ApiExecption {
	uintID, err := strconv.ParseUint(bookID, 10, 0)

	if err != nil {
		return  execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	err = b.bookRepo.Delete(uint(uintID))

	if err != nil {
		if err == repository.ErrCBookNotFound {
			return execption.NewApiExecption(http.StatusNotFound, err.Error())
		}
		return  execption.NewApiExecption(http.StatusInternalServerError, err.Error())
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

	response := &dto.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Price:       book.Price,
		ImagePath:   book.ImagePath,
		Description: book.Description,
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
		responses = append(responses, &dto.BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			Price:       book.Price,
			ImagePath:   book.ImagePath,
			Description: book.Description,
			CreatedAt:   book.CreatedAt.String(),
			UpdatedAt:   book.UpdatedAt.String(),
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

	price, err := strconv.Atoi(input.Price)

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusBadRequest, "Invalid price format")
	}

	book := &entity.Book{
		ID:   uint(bookID),
		Title:     input.Title,
		Price:     price,
		ImagePath: input.ImagePath,
		Description: input.Description,
	}

	book, err = b.bookRepo.Update(book)

	if err != nil {
		if err == repository.ErrCBookNotFound {
			return nil, execption.NewApiExecption(http.StatusNotFound, err.Error())
		}
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	response := &dto.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Price:       book.Price,
		ImagePath:   book.ImagePath,
		Description: book.Description,
		CreatedAt:   book.CreatedAt.String(),
		UpdatedAt:   book.UpdatedAt.String(),
	}

	return response, nil
}
