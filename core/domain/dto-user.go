package domain

import (
	"hexagonal-fiber-impl/common"
)

type UserLoginRequest struct {
	Email    string `validate:"required,min=3,max=32"`
	Password string `validate:"required,min=8,max=16"`
}

type UserRegisterRequest struct {
	Email           string `validate:"required,min=3,max=32"`
	Password        string `validate:"required,min=8,max=16"`
	ConfirmPassword string `validate:"required,min=8,max=16"`
}

func (ul UserLoginRequest) Validate() []*common.ErrorResponse {
	err := common.Validate.Struct(ul)
	if err != nil {
		return common.ValidationError(err)
	}
	return nil
}

func (ur UserRegisterRequest) Validate() []*common.ErrorResponse {
	err := common.Validate.Struct(ur)
	if err != nil {
		return common.ValidationError(err)
	}
	return nil
}
