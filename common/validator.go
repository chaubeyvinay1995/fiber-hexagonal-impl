package common

import "github.com/go-playground/validator/v10"

// Validator is used to validate each request's data.
//
// Every API exposed should have a validation on the data being received
type Validator interface {
	// Validate function is used to validate the request.
	//
	// This is the implementation of the Validator interface.
	// If any custom validation is required for a dto apart from the
	// validation(s) present as part of tags in the struct, this function
	// can be used to perform the same.
	//
	// Note : All the APIs which follows rest standard is forced to
	// implement this function. As a compliment for all such APIs,
	// validation will be performed automatically using the common rest handler.
	Validate() error
}

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func ValidationError(err error) []*ErrorResponse {
	var errors []*ErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		var element ErrorResponse
		element.Field = err.StructNamespace()
		element.Tag = err.Tag()
		element.Value = err.Param()
		errors = append(errors, &element)
	}
	return errors
}

//var validate = validator.New()

//type ErrorResponse struct {
//	Field string `json:"field"`
//	Tag   string `json:"tag"`
//	Value string `json:"value,omitempty"`
//}

//func ValidateStruct[T any](payload T) []*ErrorResponse {
//	var errors []*ErrorResponse
//	err := validate.Struct(payload)
//	if err != nil {
//		for _, err := range err.(validator.ValidationErrors) {
//			var element ErrorResponse
//			element.Field = err.StructNamespace()
//			element.Tag = err.Tag()
//			element.Value = err.Param()
//			errors = append(errors, &element)
//		}
//	}
//	return errors
//}
