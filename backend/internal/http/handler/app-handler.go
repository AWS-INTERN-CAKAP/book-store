package handler

import "github.com/aws-cakap-intern/book-store/pkg/validator"

type AppHandler struct {
	CategoryHandler *CategotyHandler
}

func NewAppHandler(categoryHandler *CategotyHandler) AppHandler {
	return AppHandler{CategoryHandler: categoryHandler}
}

func checkValidation(input interface{}) (errorMessage string, data interface{}) {
	validationErrors := validator.Validate(input)
	if validationErrors != nil {
		if _, exists := validationErrors["error"]; exists {
			return "validasi input gagal", nil
		}
		return "validasi input gagal", validationErrors
	}
	return "", nil
}