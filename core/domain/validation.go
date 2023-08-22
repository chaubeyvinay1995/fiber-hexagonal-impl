package domain

import (
	"github.com/go-playground/validator/v10"
	"hexagonal-fiber-impl/common"
)

type UserLoginRequest struct {
	Name  string `validate:"required,min=3,max=32"`
	Email string `validate:"required,email,min=6,max=32"`
}

//var validate = validator.New()

type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

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

func (ca UserLoginRequest) Validate() []*ErrorResponse {
	var errors []*ErrorResponse
	err := common.Validate.Struct(ca)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
