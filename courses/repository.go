package courses

import "gorm.io/gorm"

type Repository interface {
	Create(course Courses) (Courses, error)
}

type repository struct {
	db *gorm.DB
}

func Repositories(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(course Courses) (Courses, error) {

	err := r.db.Create(&course).Error

	if err != nil {
		return Courses{}, err
	}

	return course, nil
}
