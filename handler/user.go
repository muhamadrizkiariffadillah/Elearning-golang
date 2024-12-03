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
		response := helper.APIResponse(fiber.StatusUnprocessableEntity, "failed", "fail to capture the request", err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(response)
	}

	validate := validator.New()
	err = validate.Struct(input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMsg := fiber.Map{
			"error": errors,
		}
		response := helper.APIResponse(fiber.StatusUnprocessableEntity, "failed", "fail to capture the parameters", errorMsg)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	newUser, err := h.service.CreateUser(input)
	if err != nil {
		errMsg := fiber.Map{
			"error": err,
		}
		response := helper.APIResponse(fiber.StatusNotFound, "failed", "fail to create user", errMsg)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse(fiber.StatusOK, "success", "successfully to create user", newUser)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *userHandler) LoginUser(c *fiber.Ctx) error {

	var input users.LoginUserInput

	err := c.BodyParser(&input)

	if err != nil {
		response := helper.APIResponse(fiber.StatusUnprocessableEntity, "failed", "fail to capture the request", err)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	validate := validator.New()

	err = validate.Struct(input)

	if err != nil {
		errors := helper.FormatError(err)
		response := helper.APIResponse(fiber.StatusUnprocessableEntity, "failed", "fail to capture the parameters", errors)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	user, err := h.service.LoginUser(input)

	if err != nil {
		response := helper.APIResponse(fiber.StatusInternalServerError, "failed", "fail to get the user", err)
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	response := helper.APIResponse(fiber.StatusOK, "success", "success to login", user)
	return c.Status(fiber.StatusOK).JSON(response)
}
