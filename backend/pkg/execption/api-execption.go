package execption

type ApiExecption struct {
	Status  int
	Message string
}

func NewApiExecption(status int, message string) *ApiExecption {
	return &ApiExecption{
		Status:  status,
		Message: message,
	}
}