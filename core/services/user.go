package services

import (
	"errors"
	"hexagonal-fiber-impl/core/domain"
	"hexagonal-fiber-impl/core/ports"
)

type UserService struct {
	userRepository ports.IUserRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IUserService = (*UserService)(nil)

func NewUserService(repository ports.IUserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

func (s *UserService) Login(email string, password string) (domain.User, error) {
	user, err := s.userRepository.Login(email, password)
	if err != nil {
		return domain.User{}, err
	}
	return user, err
}

func (s *UserService) Register(email string, password string, confirmPass string) (domain.User, error) {
	if password != confirmPass {
		return domain.User{}, errors.New("password and confirm password are not same")
	}
	user, err := s.userRepository.Register(email, password)
	if err != nil {
		return domain.User{}, err
	}
	return user, err
}
