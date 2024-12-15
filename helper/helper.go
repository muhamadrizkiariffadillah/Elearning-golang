package helper

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

func FormatError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}

func IsEmail(input string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(input)
}

type Reposnse struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func APIResponse(code int, status, message string, data any) Reposnse {

	format := Reposnse{
		Meta: Meta{
			Code:    code,
			Status:  status,
			Message: message,
		},
		Data: data,
	}

	return format
}

type UserFormat struct {
	Fullname     string `json:"full_name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	IsMembership bool   `json:"membership"`
}

func UserFormatter(fullname, username, email string, membership time.Time) UserFormat {

	formatter := UserFormat{
		Fullname: fullname,
		Username: username,
		Email:    email,
	}

	if membership.Unix() <= time.Now().Unix() {
		formatter.IsMembership = false
	}

	return formatter
}

type CourseFormat struct {
	CourseName       string `json:"course_name"`
	CourseImageUrl   string `json:"course_img_url"`
	ShortDescription string `json:"short_description"`
	Price            int    `json:"price"`
	Discount         uint8  `json:"discount"`
	FinalPrice       int    `json:"final_price"`
}

func CourseFormatter(name, imageUrl, sd string, p, d, fp int) CourseFormat {

	formatter := CourseFormat{
		CourseName:       name,
		CourseImageUrl:   imageUrl,
		ShortDescription: sd,
		Price:            p,
		Discount:         uint8(d),
		FinalPrice:       fp,
	}

	return formatter

}
