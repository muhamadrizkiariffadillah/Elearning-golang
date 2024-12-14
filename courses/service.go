package courses

import (
	"errors"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateCourse(input CreateCourseInput) (Courses, error)
	UpdateCourse(id int, input CreateCourseInput) (Courses, error)
}

type service struct {
	repo Repository
}

func Services(r Repository) *service {
	return &service{r}
}

func (s *service) CreateCourse(input CreateCourseInput) (Courses, error) {

	newCourse := Courses{
		CourseName:       input.CourseName,
		CourseImageUrl:   input.CourseImageUrl,
		ShortDescription: input.ShortDescription,
		Price:            input.Price,
		DiscountPercent:  input.DiscountPrice,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	slug := slug.Make(input.CourseName)

	newCourse.Slug = slug

	discount := int((input.Price * int(input.DiscountPrice)) / 100)

	newCourse.FinalPrice = int(input.Price - discount)

	course, err := s.repo.Create(newCourse)

	if err != nil {
		return Courses{}, err
	}

	return course, nil
}

func (s *service) UpdateCourse(id int, input CreateCourseInput) (Courses, error) {

	course, err := s.repo.FindById(id)

	if err != nil {
		return Courses{}, err
	}

	if course.Id == 0 {
		return Courses{}, errors.New("course not found")
	}

	course.CourseName = input.CourseName

	course.CourseImageUrl = input.CourseImageUrl

	course.ShortDescription = input.ShortDescription

	course.DiscountPercent = input.DiscountPrice

	course.Price = input.Price

	discount := int((input.Price * int(input.DiscountPrice)) / 100)

	course.FinalPrice = int(input.Price - discount)

	updatedCourse, err := s.repo.Update(course)

	if err != nil {
		return Courses{}, err
	}

	return updatedCourse, nil

}
