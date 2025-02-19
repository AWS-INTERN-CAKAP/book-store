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
	bookRepository := repository.NewBookRepository(db)

	categoryService := service.NewCategoryService(categoryRepository)
	bookService := service.NewBookService(bookRepository, categoryRepository)
	
	categoryHandler := handler.NewCategoryHandler(categoryService)
	bookHandler := handler.NewBookHandler(bookService)

	appHandler := handler.NewAppHandler(categoryHandler, bookHandler)

	return router.AppPublicRoutes(appHandler)
}
