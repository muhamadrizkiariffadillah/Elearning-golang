package handler

import (
	"e-learning/courses"
	"e-learning/helper"
	"strconv"

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

	formatter := helper.CourseFormatter(newCourse.CourseName, newCourse.CourseImageUrl, newCourse.ShortDescription, newCourse.Price, int(newCourse.DiscountPercent), newCourse.FinalPrice)

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

	formatter := helper.CourseFormatter(updatedCourse.CourseName, updatedCourse.CourseImageUrl, updatedCourse.ShortDescription, updatedCourse.Price, int(updatedCourse.DiscountPercent), updatedCourse.FinalPrice)

	response := helper.APIResponse(fiber.StatusOK, "success", "success to update the course", formatter)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *courseHandler) GetCourseById(c *fiber.Ctx) error {

	var input courses.Param

	err := c.ParamsParser(&input)

	if err != nil {
		errorMsg := fiber.Map{
			"error": err.Error(),
		}

		response := helper.APIResponse(fiber.StatusBadRequest, "failed", "fail to capture the uri", errorMsg)

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	course, err := h.service.FindCourseById(input.Id)

	if err != nil {
		errorMsg := fiber.Map{"error": err.Error()}

		response := helper.APIResponse(fiber.StatusInternalServerError, "failed", "fail to get the course", errorMsg)

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	formatter := helper.CourseFormatter(course.CourseName, course.CourseImageUrl, course.ShortDescription, course.Price, int(course.DiscountPercent), course.FinalPrice)

	response := helper.APIResponse(fiber.StatusOK, "success", "success to update the course", formatter)

	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *courseHandler) CreateSubCourseByCourseId(c *fiber.Ctx) error {

	var param courses.Param
	var input courses.CreateSubCourseInput

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

		response := helper.APIResponse(fiber.StatusBadRequest, "failed", "fail to capture the uri", errorMsg)

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	validate := validator.New()

	err = validate.Struct(input)
	if err != nil {
		fmtError := helper.FormatError(err)
		errorMsg := fiber.Map{
			"error": fmtError,
		}

		response := helper.APIResponse(fiber.StatusBadRequest, "failed", "fail to capture the params", errorMsg)

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	newSubCourse, err := h.service.CreateSubCourse(param.Id, input)

	if err != nil {
		errorMsg := fiber.Map{
			"error": err.Error(),
		}

		response := helper.APIResponse(fiber.StatusBadRequest, "failed", "fail to capture the params", errorMsg)

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	response := helper.APIResponse(fiber.StatusOK, "success", "success to create a subcourse", newSubCourse)

	return c.Status(fiber.StatusOK).JSON(response)

}

func (h *courseHandler) UpdateSubCourse(c *fiber.Ctx) error {

	var params courses.UpdateSubParams

	id := c.Params("id")
	subId := c.Params("sub_id")

	params.CourseId, _ = strconv.Atoi(id)
	params.SubCourseId, _ = strconv.Atoi(subId)

	var input courses.CreateSubCourseInput

	err := c.ParamsParser(&params)

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

		response := helper.APIResponse(fiber.StatusBadRequest, "failed", "fail to capture the uri", errorMsg)

		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	validate := validator.New()

	err = validate.Struct(params)

	if err != nil {

		fmtError := helper.FormatError(err)
		errorMsg := fiber.Map{
			"error": fmtError,
		}

		response := helper.APIResponse(fiber.StatusUnprocessableEntity, "failed", "fail to capture the params", errorMsg)

		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	err = validate.Struct(input)

	if err != nil {

		fmtError := helper.FormatError(err)
		errorMsg := fiber.Map{
			"error": fmtError,
		}

		response := helper.APIResponse(fiber.StatusUnprocessableEntity, "failed", "fail to capture the params", errorMsg)

		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	updateSubCourse, err := h.service.UpdateSubCourse(params, input)

	if err != nil {
		response := helper.APIResponse(fiber.StatusInternalServerError, "failed", "fail to capture the params", fiber.Map{"error": err.Error()})

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse(fiber.StatusOK, "success", "success to update the sub course.", fiber.Map{"sub_course": updateSubCourse})

	return c.Status(fiber.StatusOK).JSON(response)
}
