package courses

type CreateCourseInput struct {
	CourseName       string `json:"course_name" validate:"required"`
	CourseImageUrl   string `json:"course_image_url" validate:"required"`
	ShortDescription string `json:"short_description" validate:"required"`
	Price            int    `json:"price" validate:"required"`
	DiscountPrice    uint8  `json:"discount" validate:"required"`
}
