package users

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(input RegisterUserInput) (Users, error)
}

type service struct {
	repository Repository
}

func Services(repo Repository) *service {
	return &service{repo}
}

func (s *service) CreateUser(input RegisterUserInput) (Users, error) {

	candidateUser := Users{
		FullName:  input.Fulname,
		Username:  input.Username,
		Email:     input.Email,
		Role:      "student",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return Users{}, err
	}

	candidateUser.HashPassword = string(hashPassword)

	newUser, err := s.repository.Create(candidateUser)

	if err != nil {
		return Users{}, err
	}

	return newUser, nil
}
