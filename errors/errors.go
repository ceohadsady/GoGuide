package errors

import (
	"net/http"
	"regexp/syntax"
)

type AppErrors struct {
	Code    int
	Message string
	Codes   syntax.ErrorCode
}

func (e AppErrors) Error() string {
	return e.Message
}

func NewNotFoundError(message string) error {
	return AppErrors{
		Code:    http.StatusNotFound,
		Message: message,
	}
}
func NewSuccessOK(message string) error {
	return AppErrors{
		Code:    http.StatusOK,
		Message: message,
	}
}
func NewSuccessCreated(message string) error {
	return AppErrors{
		Code:    http.StatusCreated,
		Message: message,
	}
}
func ErrorBadRequest(errorMessage string) error {
	return AppErrors{
		Code:    http.StatusBadRequest,
		Message: errorMessage,
	}
}
func ErrorUnprocessableEntity(errorMessage string) error {
	return AppErrors{
		Code:    http.StatusUnprocessableEntity,
		Message: errorMessage,
	}
}
