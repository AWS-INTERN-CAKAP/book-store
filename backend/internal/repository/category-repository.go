package repository

import (
	"errors"

	"github.com/aws-cakap-intern/book-store/internal/entity"
	"gorm.io/gorm"
)

var ErrCategoryNotFound = errors.New("category not found")

type CategoryRepository interface {
	Create(category *entity.Category) (*entity.Category, error)
	Update(category *entity.Category) (*entity.Category, error)
	Delete(id uint) error
	GetAll() ([]entity.Category, error)
	GetById(id uint) (*entity.Category, error)
	FindByIDs(ids []uint, categories *[]*entity.Category) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Create(category *entity.Category) (*entity.Category, error) {
	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) Update(category *entity.Category) (*entity.Category, error) {
	var existingCategory entity.Category
	if err := r.db.First(&existingCategory, category.ID).Error; err != nil {
		return nil, ErrCategoryNotFound
	}

	if err := r.db.Model(&existingCategory).Updates(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) Delete(id uint) error {
	if err := r.db.Delete(&entity.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *categoryRepository) GetAll() ([]entity.Category, error) {
	var categories []entity.Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) GetById(id uint) (*entity.Category, error) {
	var category entity.Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, ErrCategoryNotFound
	}
	return &category, nil
}

func (c *categoryRepository) FindByIDs(ids []uint, categories *[]*entity.Category) error {
	if err := c.db.Where("id IN (?)", ids).Find(categories).Error; err != nil {
		return err
	}
	return nil
}
