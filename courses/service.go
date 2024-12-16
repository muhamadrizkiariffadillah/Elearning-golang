package courses

import (
	"errors"
	"fmt"
	"time"

	"github.com/gosimple/slug"
)

type Service interface {
	CreateCourse(input CreateCourseInput) (Courses, error)
	UpdateCourse(id int, input CreateCourseInput) (Courses, error)
	FindCourseById(id int) (Courses, error)
	CreateSubCourse(id int, input CreateSubCourseInput) (SubCourses, error)
	UpdateSubCourse(params UpdateSubParams, input CreateSubCourseInput) (SubCourses, error)
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

func (s *service) FindCourseById(id int) (Courses, error) {

	course, err := s.repo.FindById(id)

	if err != nil {
		return Courses{}, errors.New("error when get the course")
	}

	if course.Id == 0 {
		return Courses{}, errors.New("the course is not found")
	}
	return course, nil
}

func (s *service) CreateSubCourse(id int, input CreateSubCourseInput) (SubCourses, error) {

	candidateSubCourse := SubCourses{
		CourseId:       id,
		SubCourseTitle: input.SubCourseTitle,
		MetadataUrl:    input.MetadataUrl,
		Description:    input.Description,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	newSubCourse, err := s.repo.CreateSub(candidateSubCourse)

	if err != nil {
		return SubCourses{}, errors.New("service: fail to create a sub course")
	}

	return newSubCourse, nil
}

func (s *service) UpdateSubCourse(params UpdateSubParams, input CreateSubCourseInput) (SubCourses, error) {

	course, err := s.repo.FindById(params.CourseId)

	fmt.Println(course.Id)

	if err != nil {
		return SubCourses{}, errors.New("error to find course id")
	}

	if course.Id == 0 {
		return SubCourses{}, errors.New("course is not found")
	}

	subCourse, err := s.repo.FindSubById(params.SubCourseId)

	if err != nil {
		return SubCourses{}, errors.New("subcourse is not found")
	}

	if subCourse.Id == 0 {
		return SubCourses{}, errors.New("subcourse is not found")
	}

	if course.Id != subCourse.CourseId {
		return SubCourses{}, errors.New("course id is not same")
	}

	subCourse.SubCourseTitle = input.SubCourseTitle
	subCourse.MetadataUrl = input.MetadataUrl
	subCourse.Description = input.Description
	subCourse.UpdatedAt = time.Now()

	updatedSubCourse, err := s.repo.UpdateSub(subCourse)

	if err != nil {
		return SubCourses{}, errors.New("error when update subcourse in service")
	}

	return updatedSubCourse, nil
}
