package users

import (
	"e-learning/helper"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(input RegisterUserInput) (Users, error)
	LoginUser(input LoginUserInput) (Users, error)
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

func (s *service) LoginUser(input LoginUserInput) (Users, error) {
	var user Users
	var err error

	inputEmailOrUsername := input.Identifier
	inputPassword := input.Password

	isEmail := helper.IsEmail(inputEmailOrUsername)

	if isEmail {
		user, err = s.repository.FindByEmail(inputEmailOrUsername)
		if err != nil {
			return Users{}, err
		}
	} else {
		user, err = s.repository.FindByUsername(inputEmailOrUsername)
		if err != nil {
			return Users{}, err
		}
	}

	if user.Id == 0 {
		return Users{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(inputPassword))
	if err != nil {
		return Users{}, errors.New("password is invalid")
	}

	return user, nil
}
