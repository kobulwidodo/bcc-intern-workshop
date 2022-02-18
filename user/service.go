package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
	LoginUser(input LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RegisterUser(input RegisterInput) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return User{}, err
	}

	user := User{
		Name:     input.Name,
		Username: input.Username,
		Password: string(passwordHash),
		Bio:      input.Bio,
	}

	registerdUser, err := s.repository.Save(user)
	if err != nil {
		return registerdUser, err
	}

	return registerdUser, nil
}

func (s *service) LoginUser(input LoginInput) (User, error) {
	var user User
	user, err := s.repository.FindByUsername(input.Username)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return user, errors.New("password tidak sesuai")
	}

	return user, nil
}
