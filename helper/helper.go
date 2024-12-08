package helper

import (
	"regexp"

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
