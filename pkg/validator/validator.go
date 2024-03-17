package validates

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type JSONValidator struct {
	validate *validator.Validate
}

func NewValidator(validate *validator.Validate) *JSONValidator {

	return &JSONValidator{
		validate: validate,
	}

}

func (cv *JSONValidator) Validate(ctx context.Context, i interface{}) error {
	// if err := cv.Validator.Struct(i); err != nil {
	// 	// Handle validation errors with custom message
	// 	var validationErrors fiber.Map
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		validationErrors[err.Field()] = err.Error()
	// 	}
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		"error":   "Validation failed",
	// 		"details": validationErrors,
	// 	})
	// }
	// return nil

	// if err := cv.Validator.Struct(i); err != nil {
	// 	c.Status(400)
	// 	return c.JSON(err.Error())
	// }
	// return nil

	// Register the custom validation function
	cv.validate.RegisterValidation("string", validateString)

	var errStr []string
	err := cv.validate.StructCtx(ctx, i)
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Error())
		errStr = append(errStr, translatedErr.Error())
	}
	errMsg := strings.Join(errStr, ", ")
	return errors.New(errMsg)
}

// Define a custom validation function
func validateString(fl validator.FieldLevel) bool {
	// Get the field value
	fieldValue := fl.Field().Interface()

	// Check if the field value is a string
	_, ok := fieldValue.(string)
	return ok
}
