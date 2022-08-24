package validation

import "github.com/go-playground/validator/v10"

type ErrorsResponse struct {
	Error string
}

func ValidateStruct(request interface{}) []*ErrorsResponse {
	var errors []*ErrorsResponse
	var validate = validator.New()
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorsResponse
			params := ""
			if err != nil {
				params = " " + err.Param()
			}
			element.Error = err.StructField() + " " + err.Tag() + params
			errors = append(errors, &element)
		}
	}
	return errors
}
