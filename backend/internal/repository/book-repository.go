package repository

import (
	"errors"

	"github.com/aws-cakap-intern/book-store/internal/entity"
	"gorm.io/gorm"
)

var ErrBookNotFound = errors.New("book not found")

type BookRepository interface {
	Create(book *entity.Book, categoryIDs []uint) (*entity.Book, error)
	Update(book *entity.Book, categoryIDs []uint) (*entity.Book, error)
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
func (b *bookRepository) Create(book *entity.Book, categoryIDs []uint) (*entity.Book, error) {
	if err := b.db.Create(book).Error; err != nil {
		return nil, err
	}

	// Associate categories
	if len(categoryIDs) > 0 {
		var categories []entity.Category
		if err := b.db.Where("id IN ?", categoryIDs).Find(&categories).Error; err != nil {
			return nil, err
		}
		b.db.Model(book).Association("Categories").Replace(categories)
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
	if err := b.db.Preload("Categories").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetById implements BookRepository.
func (b *bookRepository) GetById(id uint) (*entity.Book, error) {
	var book entity.Book
	if err := b.db.Preload("Categories").First(&book, id).Error; err != nil {
		return nil, ErrBookNotFound
	}
	return &book, nil
}

// Update implements BookRepository.
func (b *bookRepository) Update(book *entity.Book, categoryIDs []uint) (*entity.Book, error) {
	var existingBook entity.Book
	if err := b.db.First(&existingBook, book.ID).Error; err != nil {
		return nil, ErrBookNotFound
	}

	// Update book details
	if err := b.db.Model(&existingBook).Updates(book).Error; err != nil {
		return nil, err
	}

	// Update category relationships
	if len(categoryIDs) > 0 {
		var categories []entity.Category
		if err := b.db.Where("id IN ?", categoryIDs).Find(&categories).Error; err != nil {
			return nil, err
		}
		b.db.Model(&existingBook).Association("Categories").Replace(categories)
	}

	return &existingBook, nil
}
