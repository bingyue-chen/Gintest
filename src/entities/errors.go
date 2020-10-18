package entities

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"io"
)

type ErrorBag struct {
	Message string `json:"message"`
}

type UnauthorizedError struct {
	Message string `json:"message"`
}

func (err *UnauthorizedError) Error() string {
	return err.Message
}

func (err *UnauthorizedError) Is(target error) bool {
	_, ok := target.(*UnauthorizedError)

	if !ok {
		return false
	}

	return true
}

type ForbiddenError struct {
	Message string `json:"message"`
}

func (err *ForbiddenError) Error() string {
	return err.Message
}

func (err *ForbiddenError) Is(target error) bool {
	_, ok := target.(*ForbiddenError)

	if !ok {
		return false
	}

	return true
}

type BindJSONError struct {
	validationErrors validator.ValidationErrors
	otherError       error
}

func (err *BindJSONError) SetErrors(errors error) {

	var ok bool

	if err.validationErrors, ok = errors.(validator.ValidationErrors); !ok {
		err.otherError = errors
	}
}

func (err *BindJSONError) Error() string {

	if err.otherError != nil {

		if errors.Is(err.otherError, io.EOF) {
			return "Body of request is required json format"
		}

		return err.otherError.Error()
	}

	for _, fieldErr := range err.validationErrors {
		return fmt.Sprintf("Field validation for '%s' failed on the '%s' rule", fieldErr.Field(), fieldErr.Tag())
	}

	return "Missing or wrong request data"
}

func (err *BindJSONError) Is(target error) bool {
	_, ok := target.(*BindJSONError)

	if !ok {
		return false
	}

	return true
}

type BadRequestError struct {
	Message string `json:"message"`
}

func (err *BadRequestError) Error() string {
	return err.Message
}

func (err *BadRequestError) Is(target error) bool {
	_, ok := target.(*BadRequestError)

	if !ok {
		return false
	}

	return true
}
