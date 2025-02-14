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

type CategoryService interface {
	GetCategories() ([]*dto.CategoryResponse, *execption.ApiExecption)
	GetCategory(categoryID string) (*dto.CategoryResponse, *execption.ApiExecption)
	CreateCategory(input binder.CreateCategory) (*dto.CategoryResponse, *execption.ApiExecption)
	UpdateCategory(input binder.UpdateCategory) (*dto.CategoryResponse, *execption.ApiExecption)
	DeleteCategory(categoryID string) *execption.ApiExecption
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (s *categoryService) GetCategories() ([]*dto.CategoryResponse, *execption.ApiExecption)  {
	categories, err := s.categoryRepo.GetAll()

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	if len(categories) == 0 {
		return []*dto.CategoryResponse{}, nil
	}

	var responses []*dto.CategoryResponse

	for _, category := range categories {
		responses = append(responses, &dto.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt.String(),
			UpdatedAt: category.UpdatedAt.String(),
		})
	}

	return responses, nil
}

func (s *categoryService) GetCategory(categoryID string) (*dto.CategoryResponse, *execption.ApiExecption)  {
	uintID, err := strconv.ParseUint(categoryID, 10, 0)

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	category, err := s.categoryRepo.GetById(uint(uintID))

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusNotFound, err.Error())
	}

	response := &dto.CategoryResponse{
		ID:        category.ID,	
		Name:      category.Name,
		CreatedAt: category.CreatedAt.String(),
		UpdatedAt: category.UpdatedAt.String(),
	}

	return response, nil
}

func (s *categoryService) CreateCategory(input binder.CreateCategory) (*dto.CategoryResponse, *execption.ApiExecption)  {
	category := &entity.Category{
		Name: input.Name,
	}

	category, err := s.categoryRepo.Create(category)

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	response := &dto.CategoryResponse{
		ID:        category.ID,	
		Name:      category.Name,
		CreatedAt: category.CreatedAt.String(),
		UpdatedAt: category.UpdatedAt.String(),
	}

	return response, nil
}

func (s *categoryService) UpdateCategory(input binder.UpdateCategory) (*dto.CategoryResponse, *execption.ApiExecption)  {
	categoryID, err := strconv.ParseUint(input.ID, 10, 0)

	if err != nil {
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	category := &entity.Category{
		ID:   uint(categoryID),
		Name: input.Name,
	}

	category, err = s.categoryRepo.Update(category)

	if err != nil {
		if err == repository.ErrCategoryNotFound {
			return nil, execption.NewApiExecption(http.StatusNotFound, err.Error())
		}
		return nil, execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	response := &dto.CategoryResponse{
		ID:        category.ID,	
		Name:      category.Name,
		CreatedAt: category.CreatedAt.String(),
		UpdatedAt: category.UpdatedAt.String(),
	}

	return response, nil
}

func (s *categoryService) DeleteCategory(categoryID string)  *execption.ApiExecption  {
	uintID, err := strconv.ParseUint(categoryID, 10, 0)

	if err != nil {
		return  execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	err = s.categoryRepo.Delete(uint(uintID))

	if err != nil {
		if err == repository.ErrCategoryNotFound {
			return execption.NewApiExecption(http.StatusNotFound, err.Error())
		}
		return  execption.NewApiExecption(http.StatusInternalServerError, err.Error())
	}

	return nil
}