package users

import "gorm.io/gorm"

type Repository interface {
	Create(user Users) (Users, error)
	FindByEmail(email string) (Users, error)
	FindByUsername(username string) (Users, error)
}

type repository struct {
	db *gorm.DB
}

func Repositories(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user Users) (Users, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return Users{}, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (Users, error) {
	var user Users

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return Users{}, err
	}
	return user, nil
}

func (r *repository) FindByUsername(username string) (Users, error) {

	var user Users

	err := r.db.Where("username = ?", username).Find(&user).Error

	if err != nil {

		return Users{}, err

	}

	return user, nil
}
