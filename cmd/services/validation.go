package services

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type Validation struct{}

func (v Validation) RequestBody(body io.ReadCloser, shape any) error {
	defer body.Close()

	decoder := json.NewDecoder(body)
	if err := decoder.Decode(shape); err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(shape); err != nil {

		if invErr, ok := err.(*validator.InvalidValidationError); ok {
			return invErr
		}

		if valErrs, ok := err.(validator.ValidationErrors); ok {
			return valErrs
		}

		return err
	}

	return nil
}
