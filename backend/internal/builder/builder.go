package builder

import (
	"github.com/aws-cakap-intern/book-store/internal/http/handler"
	"github.com/aws-cakap-intern/book-store/internal/http/router"
	"github.com/aws-cakap-intern/book-store/internal/repository"
	"github.com/aws-cakap-intern/book-store/internal/service"
	"github.com/aws-cakap-intern/book-store/pkg/route"
	"gorm.io/gorm"
)

func BuildAppPublicRoutes(db *gorm.DB, ) []*route.Route {
	

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)


	appHandler := handler.NewAppHandler(categoryHandler)

	return router.AppPublicRoutes(appHandler)
}
