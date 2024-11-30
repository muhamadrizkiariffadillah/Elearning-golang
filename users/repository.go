package users

import "gorm.io/gorm"

type Repository interface {
	Create(user Users) (Users, error)
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
