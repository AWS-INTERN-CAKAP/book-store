package handler

import (
	"net/http"

	"github.com/aws-cakap-intern/book-store/internal/http/binder"
	"github.com/aws-cakap-intern/book-store/internal/service"
	"github.com/aws-cakap-intern/book-store/pkg/response"
	"github.com/labstack/echo/v4"
)

type CategotyHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) *CategotyHandler {
	return &CategotyHandler{categoryService: categoryService}
}

func (c *CategotyHandler) GetCategories(ctx echo.Context) error {
	responsData, execption := c.categoryService.GetCategories()

	if execption != nil {
		return ctx.JSON(execption.Status, response.ErrorResponse(execption.Status, execption.Message))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success Get Categories", responsData))
}

func (c *CategotyHandler) GetCategory(ctx echo.Context) error {
	var input binder.GetCategory

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	
	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	responsData, execption := c.categoryService.GetCategory(input.ID)

	if execption != nil {
		return ctx.JSON(execption.Status, response.ErrorResponse(execption.Status, execption.Message))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success Get Category", responsData))
}

func (c *CategotyHandler) CreateCategory(ctx echo.Context) error {
	var input binder.CreateCategory

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	
	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	responsData, execption := c.categoryService.CreateCategory(input)

	if execption != nil {
		return ctx.JSON(execption.Status, response.ErrorResponse(execption.Status, execption.Message))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success Create Category", responsData))
}

func (c *CategotyHandler) UpdateCategory(ctx echo.Context) error {
	var input binder.UpdateCategory

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	
	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	responsData, execption := c.categoryService.UpdateCategory(input)

	if execption != nil {
		return ctx.JSON(execption.Status, response.ErrorResponse(execption.Status, execption.Message))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success Update Category", responsData))
}

func (c *CategotyHandler) DeleteCategory(ctx echo.Context) error {
	var input binder.DeleteCategory

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	
	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	execption := c.categoryService.DeleteCategory(input.ID)

	if execption != nil {
		return ctx.JSON(execption.Status, response.ErrorResponse(execption.Status, execption.Message))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success Delete Category", nil))
}