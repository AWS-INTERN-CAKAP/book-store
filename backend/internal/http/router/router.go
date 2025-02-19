package router

import (
	"net/http"

	"github.com/aws-cakap-intern/book-store/internal/http/handler"
	"github.com/aws-cakap-intern/book-store/pkg/route"
	"github.com/labstack/echo/v4"
)

func AppPublicRoutes(appHandler handler.AppHandler) []*route.Route {
	categoryHandler := appHandler.CategoryHandler
	bookHandler := appHandler.BookHandler

	return []*route.Route{
		{
			Method:  http.MethodGet,
			Path:    "/categories",
			Handler: categoryHandler.GetCategories,
		},
		{
			Method:  http.MethodGet,
			Path:    "/categories/:id",
			Handler: categoryHandler.GetCategory,
		},
		{
			Method:  http.MethodPost,
			Path:    "/categories",
			Handler: categoryHandler.CreateCategory,
		},
		{
			Method:  http.MethodPut,
			Path:    "/categories/:id",
			Handler: categoryHandler.UpdateCategory,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/categories/:id",
			Handler: categoryHandler.DeleteCategory,
		},
		{
			Method:  http.MethodGet,
			Path:    "/books",
			Handler: bookHandler.GetBooks,
		},
		{
			Method:  http.MethodGet,
			Path:    "/books/:id",
			Handler: bookHandler.GetBook,
		},
		{
			Method:  http.MethodPost,
			Path:    "/books",
			Handler: bookHandler.CreateBook,
		},
		{
			Method:  http.MethodPut,
			Path:    "/books/:id",
			Handler: bookHandler.UpdateBook,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/books/:id",
			Handler: bookHandler.DeleteBook,
		},
		{
			Method: http.MethodGet,
			Path:   "/uploads/*",
			Handler: echo.WrapHandler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads")))),
		},
	}
}
