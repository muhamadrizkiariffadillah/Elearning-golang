package courses

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(course Courses) (Courses, error)
	Update(course Courses) (Courses, error)
	FindById(id int) (Courses, error)
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

func (r *repository) Update(course Courses) (Courses, error) {

	err := r.db.Model(&course).Updates(Courses{
		CourseName:      course.CourseName,
		CourseImageUrl:  course.CourseImageUrl,
		Price:           course.Price,
		DiscountPercent: course.DiscountPercent,
		FinalPrice:      course.FinalPrice,
		UpdatedAt:       course.UpdatedAt,
	}).Error

	if err != nil {
		return Courses{}, err
	}

	return course, nil
}

func (r *repository) FindById(id int) (Courses, error) {

	var course Courses

	err := r.db.Where("id = ?", id).Find(&course).Error

	if err != nil {
		return Courses{}, err
	}

	return course, nil
}
