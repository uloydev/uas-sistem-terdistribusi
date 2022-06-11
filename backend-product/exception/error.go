package exception

import "fmt"

func PanicWhenError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicWhenEnvNil(key string, env interface{}) {
	if env == nil || env == "" {
		panic(fmt.Sprintf("ENV %s cannot be empty", key))
	}
}

func PanicValidationWhenError(err error) {
	if err != nil {
		panic(ValidationError{
			Message: err.Error(),
		})
	}
}
