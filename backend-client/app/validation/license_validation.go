package validation

import (
	"sister-backend/app/model"
	"sister-backend/exception"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func ValidateLicense(request model.LicenseRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Key, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
