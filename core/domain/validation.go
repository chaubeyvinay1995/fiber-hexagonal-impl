package domain

import (
	"github.com/go-playground/validator/v10"
	"hexagonal-fiber-impl/common"
)

type UserLoginRequest struct {
	Email    string `validate:"required,min=3,max=32"`
	Password string `validate:"required,min=8,max=16"`
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
