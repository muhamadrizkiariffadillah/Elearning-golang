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
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
