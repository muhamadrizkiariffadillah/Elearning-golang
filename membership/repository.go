package membership

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(entity Membership) (Membership, error)
	Update(entity Membership) (Membership, error)
	FindByUserId(id int) (Membership, error)
}

type repository struct {
	db *gorm.DB
}

func Repositories(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(entity Membership) (Membership, error) {

	err := r.db.Create(&entity).Error
	if err != nil {
		return Membership{}, err
	}

	return entity, nil
}

func (r *repository) Update(entity Membership) (Membership, error) {

	err := r.db.Model(&entity).Updates(entity).Error

	if err != nil {
		return Membership{}, errors.New("error to update into database")
	}

	return entity, nil
}

func (r *repository) FindByUserId(id int) (Membership, error) {

	var membership Membership

	err := r.db.Where("id = ?", id).Find(&membership).Error

	if err != nil {
		return Membership{}, errors.New("error when find membership in database")
	}

	return membership, nil
}
