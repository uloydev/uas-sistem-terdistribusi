package validation

import (
	"sister-backend/app/model"
	"sister-backend/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateUser(request model.UserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Birthday, is.Digit),
		validation.Field(&request.Phone, is.Digit),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
