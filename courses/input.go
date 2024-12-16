package courses

type CreateCourseInput struct {
	CourseName       string `json:"course_name" validate:"required"`
	CourseImageUrl   string `json:"course_image_url" validate:"required"`
	ShortDescription string `json:"short_description" validate:"required"`
	Price            int    `json:"price" validate:"required"`
	DiscountPrice    uint8  `json:"discount" validate:"required"`
}

type Param struct {
	Id int `uri:"id"`
}

type CreateSubCourseInput struct {
	SubCourseTitle string `json:"title" validate:"required"`
	MetadataUrl    string `json:"metadata_url" validate:"required"`
	Description    string `json:"description" validate:"required"`
}

type UpdateSubParams struct {
	CourseId    int `uri:"id"`
	SubCourseId int `uri:"sub_id"`
}
