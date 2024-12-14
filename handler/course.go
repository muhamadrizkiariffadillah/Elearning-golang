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

	formatter := helper.CourseFormatter(newCourse.CourseName, newCourse.CourseImageUrl, newCourse.ShortDescription, newCourse.FinalPrice)

	response := helper.APIResponse(fiber.StatusOK, "success", "successfully to create a new course", formatter)

	return c.Status(fiber.StatusOK).JSON(response)

}

func (h *courseHandler) UpdateCourse(c *fiber.Ctx) error {

	var param courses.Param
	var input courses.CreateCourseInput

	err := c.ParamsParser(&param)

	if err != nil {
		errorMsg := fiber.Map{
			"error": err.Error(),
		}

		response := helper.APIResponse(fiber.StatusBadRequest, "failed", "fail to capture the uri", errorMsg)

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	err = c.BodyParser(&input)

	if err != nil {
		errorMsg := fiber.Map{
			"error": err.Error(),
		}

		response := helper.APIResponse(fiber.StatusBadRequest, "failed", "fail to capture the request", errorMsg)

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	validate := validator.New()

	err = validate.Struct(input)

	if err != nil {
		errorMsg := fiber.Map{"error": err.Error()}

		response := helper.APIResponse(fiber.StatusUnprocessableEntity, "failed", "fail to get param", errorMsg)

		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	updatedCourse, err := h.service.UpdateCourse(param.Id, input)

	if err != nil {
		errorMsg := fiber.Map{"error": err.Error()}

		response := helper.APIResponse(fiber.StatusInternalServerError, "failed", "fail to get param", errorMsg)

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	formatter := helper.CourseFormatter(updatedCourse.CourseName, updatedCourse.CourseImageUrl, updatedCourse.ShortDescription, updatedCourse.FinalPrice)

	response := helper.APIResponse(fiber.StatusOK, "success", "success to update the course", formatter)

	return c.Status(fiber.StatusOK).JSON(response)
}
