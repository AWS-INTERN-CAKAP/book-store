package repository

import (
	"errors"

	"github.com/aws-cakap-intern/book-store/internal/entity"
	"gorm.io/gorm"
)

var ErrCBookNotFound = errors.New("book not found")

type BookRepository interface {
	Create(book *entity.Book) (*entity.Book, error)
	Update(book *entity.Book) (*entity.Book, error)
	Delete(id uint) error
	GetAll() ([]entity.Book, error)
	GetById(id uint) (*entity.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db}
}

// Create implements BookRepository.
func (b *bookRepository) Create(book *entity.Book) (*entity.Book, error) {
	if err := b.db.Create(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

// Delete implements BookRepository.
func (b *bookRepository) Delete(id uint) error {
	if err := b.db.Delete(&entity.Book{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetAll implements BookRepository.
func (b *bookRepository) GetAll() ([]entity.Book, error) {
	var books []entity.Book
	if err := b.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetById implements BookRepository.
func (b *bookRepository) GetById(id uint) (*entity.Book, error) {
	var book entity.Book
	if err := b.db.Find(&book, id).Error; err != nil {
		return nil, ErrCBookNotFound
	}
	return &book, nil
}

// Update implements BookRepository.
func (b *bookRepository) Update(book *entity.Book) (*entity.Book, error) {
	var existingBook entity.Book
	if err := b.db.First(&existingBook, book.ID).Error; err != nil {
		return nil, ErrCBookNotFound
	}

	if err := b.db.Model(&existingBook).Updates(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}


