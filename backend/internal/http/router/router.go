package router

import (
	"net/http"

	"github.com/aws-cakap-intern/book-store/internal/http/handler"
	"github.com/aws-cakap-intern/book-store/pkg/route"
)

func AppPublicRoutes(appHandler handler.AppHandler) []*route.Route {
	categoryHandler := appHandler.CategoryHandler
	
	return []*route.Route{
			{
				Method: http.MethodGet,
				Path: "/categories",
				Handler: categoryHandler.GetCategories,
			},
			{
				Method: http.MethodGet,
				Path: "/categories/:id",
				Handler: categoryHandler.GetCategory,
			},
			{
				Method: http.MethodPost,
				Path: "/categories",
				Handler: categoryHandler.CreateCategory,
			},
			{
				Method: http.MethodPut,
				Path: "/categories/:id",
				Handler: categoryHandler.UpdateCategory,
			},
			{
				Method: http.MethodDelete,
				Path: "/categories/:id",
				Handler: categoryHandler.DeleteCategory,
			},
	}
}