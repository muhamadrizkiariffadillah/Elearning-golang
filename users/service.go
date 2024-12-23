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
	UpdateUserInfo(id int, input UpdateUserInput) (Users, error)
	UpdatePassword(id int, input UpdatePasswordInput) (Users, error)
	FindUserById(id int) (Users, error)
	FindUserByUsername(username string) (Users, error)
	CreateUserProgress(userId, courseId, subCourseId int) (UserProgesses, error)
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
		Username:  input.Username, // identifier
		Email:     input.Email,    //identifier
		Role:      "student",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// this is used for checking username

	checkNewUserUsername, err := s.repository.FindByUsername(candidateUser.Username)

	if err != nil {
		return Users{}, err
	}

	if checkNewUserUsername.Id != 0 {
		return Users{}, errors.New("username has already taken")
	}

	// this is used for checking email
	checkNewUserEmail, err := s.repository.FindByEmail(candidateUser.Email)

	if err != nil {
		return Users{}, err
	}

	if checkNewUserEmail.Id != 0 {
		return Users{}, errors.New("email has already taken")
	}

	// this is used for encrypt the password

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

func (s *service) UpdateUserInfo(id int, input UpdateUserInput) (Users, error) {

	user, err := s.repository.FindById(id)

	if err != nil {
		return Users{}, err
	}

	user.Email = input.Email

	checkUserEmail, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return Users{}, err
	}

	if checkUserEmail.Id != 0 {
		return Users{}, errors.New("email is used for another account")
	}

	user.Username = input.Username

	checkUserUsename, err := s.repository.FindByUsername(input.Username)

	if err != nil {
		return Users{}, err
	}

	if checkUserUsename.Id != 0 {
		return Users{}, errors.New("email is used for another account")
	}

	user.FullName = input.Fullname

	updatedUser, err := s.repository.UpdateInfo(user)

	if err != nil {
		return Users{}, err
	}

	return updatedUser, nil

}

func (s *service) UpdatePassword(id int, input UpdatePasswordInput) (Users, error) {

	user, err := s.repository.FindById(id)

	if err != nil {
		return Users{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(input.OldPassword))

	if err != nil {
		return user, errors.New("old password is wrong")
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)

	if err != nil {
		return user, err
	}

	user.HashPassword = string(newPassword)

	updatedUserPassword, err := s.repository.UpdatePassword(user)

	if err != nil {
		return user, err
	}

	return updatedUserPassword, nil
}

func (s *service) FindUserById(id int) (Users, error) {

	user, err := s.repository.FindById(id)

	if err != nil {
		return Users{}, nil
	}

	if user.Id == 0 {
		return Users{}, errors.New("user is not found")
	}

	return user, nil
}

func (s *service) FindUserByUsername(username string) (Users, error) {

	user, err := s.repository.FindByUsername(username)

	if err != nil {
		return Users{}, err
	}

	return user, nil
}

func (s *service) CreateUserProgress(userId, courseId, subCourseId int) (UserProgesses, error) {
	// Todo: create a solution to solve a problem when save lots of userprogresses.
	// course 1
	// 1,3,5,6,10,12
	progres := UserProgesses{
		UserId:      userId,
		CourseId:    courseId,
		SubCourseId: subCourseId,
		IsComplete:  false,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	userProgress, err := s.repository.CreateUserProgess(progres)

	if err != nil {
		return UserProgesses{}, errors.New("error service to save the progress")
	}

	return userProgress, nil

}
