package handler

import (
	"e-learning/courses"
	"e-learning/helper"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type courseHandler struct {
	service courses.Service
}

func CourseHandlers(s courses.Service) *courseHandler {
	return &courseHandler{s}
}

func (h *courseHandler) CreateNewCourse(c *fiber.Ctx) error {

	var input courses.CreateCourseInput

	err := c.BodyParser(&input)

	if err != nil {

		response := helper.APIResponse(fiber.StatusBadRequest, "failed", "fail to capture the body request", err.Error())

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	validate := validator.New()

	err = validate.Struct(input)

	if err != nil {

		formatErr := helper.FormatError(err)

		response := helper.APIResponse(fiber.StatusUnprocessableEntity, "failed", "fail to get the parameters", formatErr)

		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	newCourse, err := h.service.CreateCourse(input)

	if err != nil {

		response := helper.APIResponse(fiber.StatusInternalServerError, "failed", "fail to create a new course", err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse(fiber.StatusOK, "success", "successfully to create a new course", newCourse)

	return c.Status(fiber.StatusOK).JSON(response)

}
