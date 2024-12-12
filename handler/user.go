package handler

import (
	"e-learning/auth"
	"e-learning/helper"
	"e-learning/membership"
	"e-learning/users"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	service          users.Service
	authService      auth.Service
	membershipSevice membership.Service
}

func UserHandler(service users.Service, authService auth.Service, membershipService membership.Service) *userHandler {
	return &userHandler{service, authService, membershipService}
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
		response := helper.APIResponse(fiber.StatusNotFound, "failed", "username or email has already taken", errMsg)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	membershipUser, err := h.membershipSevice.CreateMembership(newUser.Id)
	if err != nil {
		errMsg := fiber.Map{
			"error": err,
		}
		response := helper.APIResponse(fiber.StatusInternalServerError, "failed", "fail to create default membership", errMsg)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse(fiber.StatusOK, "success", "successfully to create user", membershipUser)
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

	token, err := h.authService.GenerateToken(user)
	if err != nil {
		response := helper.APIResponse(fiber.StatusInternalServerError, "failed", "fail to create user token", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse(fiber.StatusOK, "success", "success to login", token)
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *userHandler) UpdateUserInfo(c *fiber.Ctx) error {

	// state variable
	var input users.UpdateUserInput

	err := c.BodyParser(&input)

	if err != nil {
		response := helper.APIResponse(fiber.StatusBadRequest, "failed", "fail to send your request", nil)
		return c.Status(fiber.StatusBadRequest).JSON(response)
	}

	validate := validator.New()

	err = validate.Struct(input)

	if err != nil {
		errosMsg := helper.FormatError(err)
		response := helper.APIResponse(fiber.StatusUnprocessableEntity, "failed", "fail to get the parameters", errosMsg)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)
	}

	user := c.Locals("currentUser").(users.Users)

	updatedUser, err := h.service.UpdateUserInfo(user.Id, input)

	if err != nil {
		response := helper.APIResponse(fiber.StatusInternalServerError, "failed", "fail to update user", err)
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse(fiber.StatusOK, "success", "success to update user info", updatedUser)

	return c.Status(fiber.StatusOK).JSON(response)

}

func (h *userHandler) UpdateUserPassword(c *fiber.Ctx) error {

	var input users.UpdatePasswordInput

	err := c.BodyParser(&input)

	if err != nil {

		response := helper.APIResponse(fiber.StatusBadRequest, "failed", "failed to capture the request", err.Error())

		return c.Status(fiber.StatusBadRequest).JSON(response)

	}

	validate := validator.New()

	err = validate.Struct(input)

	if err != nil {

		formatError := helper.FormatError(err)

		response := helper.APIResponse(fiber.StatusUnprocessableEntity, "failed", "fail to get the parameters", formatError)

		return c.Status(fiber.StatusUnprocessableEntity).JSON(response)

	}

	user := c.Locals("currentUser").(users.Users)

	updatedUser, err := h.service.UpdatePassword(user.Id, input)

	if err != nil {

		response := helper.APIResponse(fiber.StatusInternalServerError, "failed", "fail to update the password", err.Error())

		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}

	response := helper.APIResponse(fiber.StatusOK, "success", "successfully to update user password", updatedUser)

	return c.Status(fiber.StatusOK).JSON(response)
}
