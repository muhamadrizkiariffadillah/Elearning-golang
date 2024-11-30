package handler

import (
	"e-learning/helper"
	"e-learning/users"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	service users.Service
}

func UserHandler(service users.Service) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) SignUpUser(c *fiber.Ctx) error {

	var input users.RegisterUserInput

	err := c.BodyParser(&input)
	if err != nil {
		errMsg := fiber.Map{
			"error": err,
		}
		return c.Status(fiber.StatusOK).JSON(errMsg)
	}

	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMsg := fiber.Map{
			"error": errors,
		}
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errorMsg)
	}
	newUser, err := h.service.CreateUser(input)
	if err != nil {
		errMsg := fiber.Map{
			"error": err,
		}
		return c.Status(fiber.StatusInternalServerError).JSON(errMsg)
	}

	return c.Status(fiber.StatusOK).JSON(newUser)
}
