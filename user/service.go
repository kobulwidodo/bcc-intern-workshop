package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
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
