package handler

import (
	"net/http"

	"github.com/aws-cakap-intern/book-store/internal/http/binder"
	"github.com/aws-cakap-intern/book-store/internal/service"
	"github.com/aws-cakap-intern/book-store/pkg/response"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (c *BookHandler) GetBooks(ctx echo.Context) error {
	responsData, execption := c.bookService.GetBooks()

	if execption != nil {
		return ctx.JSON(execption.Status, response.ErrorResponse(execption.Status, execption.Message))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Success Get Books", responsData))
}

func (c *BookHandler) GetBook(ctx echo.Context) error {
	var input binder.GetBook

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	
	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	responsData, execption := c.bookService.GetBook(input.ID)

	if execption != nil {
		return ctx.JSON(execption.Status, response.ErrorResponse(execption.Status, execption.Message))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Success Get Book", responsData))
}

func (c *BookHandler) CreateBook(ctx echo.Context) error {
	var input binder.CreateBook

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	
	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	responsData, execption := c.bookService.CreateBook(input)

	if execption != nil {
		return ctx.JSON(execption.Status, response.ErrorResponse(execption.Status, execption.Message))
	}

	return ctx.JSON(http.StatusCreated, response.SuccessResponse(http.StatusCreated, "Success Create Book", responsData))
}

func (c *BookHandler) UpdateBook(ctx echo.Context) error {
	var input binder.UpdateBook

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	
	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	responsData, execption := c.bookService.UpdateBook(input)

	if execption != nil {
		return ctx.JSON(execption.Status, response.ErrorResponse(execption.Status, execption.Message))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Success Update Book", responsData))
}

func (c *BookHandler) DeleteBook(ctx echo.Context) error {
	var input binder.DeleteBook

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	
	if errorMessage, data := checkValidation(input); errorMessage != "" {
		return ctx.JSON(http.StatusBadRequest, response.SuccessResponse(http.StatusBadRequest, errorMessage, data))
	}

	execption := c.bookService.DeleteBook(input.ID)

	if execption != nil {
		return ctx.JSON(execption.Status, response.ErrorResponse(execption.Status, execption.Message))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Success Delete Book", nil))
}
