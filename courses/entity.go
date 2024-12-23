package courses

import "time"

type Courses struct {
	Id               int
	CourseName       string
	CourseImageUrl   string
	ShortDescription string
	Price            int
	DiscountPercent  uint8
	FinalPrice       int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	SubCourse        []SubCourses `gorm:"foreignKey:Id"`
}

type SubCourses struct {
	Id             int
	CourseId       int
	SubCourseTitle string
	MetadataUrl    string
	Description    string
	Slug           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
