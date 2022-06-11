package validation

import (
	"sister-backend/app/model"
	"sister-backend/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func ValidateProduct(request model.ProductRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Title, validation.Required),
		validation.Field(&request.Price, validation.Required, is.Digit),
		validation.Field(&request.Description, validation.Required),
		validation.Field(&request.Stock, validation.Required, is.Digit),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
