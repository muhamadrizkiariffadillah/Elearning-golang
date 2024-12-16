package courses

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Create(course Courses) (Courses, error)
	Update(course Courses) (Courses, error)
	FindById(id int) (Courses, error)
	CreateSub(sub SubCourses) (SubCourses, error)
	UpdateSub(sub SubCourses) (SubCourses, error)
	FindSubById(id int) (SubCourses, error)
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

func (r *repository) CreateSub(sub SubCourses) (SubCourses, error) {

	err := r.db.Create(&sub).Error

	if err != nil {
		return SubCourses{}, errors.New("error when create a sub course")
	}

	return sub, nil
}

func (r *repository) UpdateSub(sub SubCourses) (SubCourses, error) {

	err := r.db.Model(&sub).Updates(sub).Error

	if err != nil {
		return SubCourses{}, errors.New("error when update to database")
	}

	return sub, nil
}

func (r *repository) FindSubById(id int) (SubCourses, error) {

	var sub SubCourses

	err := r.db.Where("id = ?", id).Find(&sub).Error

	if err != nil {
		return SubCourses{}, errors.New("error when try to search sub course in repository")
	}

	return sub, nil
}
