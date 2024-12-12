package courses

import "time"

type Service interface {
	CreateCourse(input CreateCourseInput) (Courses, error)
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

	discount := int((input.Price * int(input.DiscountPrice)) / 100)

	newCourse.FinalPrice = int(input.Price - discount)

	course, err := s.repo.Create(newCourse)

	if err != nil {
		return Courses{}, err
	}

	return course, nil
}
