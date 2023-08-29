package services

import (
	"errors"
	"fmt"
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
	fmt.Println("User Password is ", user.Password)
	if !CompareHashPassword(user.Password, password) {
		return domain.User{}, errors.New("invalid Password")
	}
	return user, err
}

func (s *UserService) Register(email string, password string, confirmPass string) (domain.User, error) {
	if password != confirmPass {
		return domain.User{}, errors.New("password and confirm password are not same")
	}
	hashPassword, err := GenerateHashPassword(password)
	if err != nil {
		return domain.User{}, err
	}
	user, err := s.userRepository.Register(email, hashPassword)
	if err != nil {
		return domain.User{}, err
	}
	return user, err
}
