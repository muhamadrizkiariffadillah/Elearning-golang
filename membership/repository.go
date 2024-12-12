package membership

import "gorm.io/gorm"

type Repository interface {
	Create(entity Membership) (Membership, error)
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
