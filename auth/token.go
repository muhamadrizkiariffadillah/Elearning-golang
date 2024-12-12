package auth

import (
	"e-learning/helper"
	"e-learning/users"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type Service interface {
	GenerateToken(user users.Users) (string, error)
	ValidateToken(token string) (*CustomClaims, error)
}

type service struct {
	secret_key []byte
}

type CustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"full_name"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func AuthServices(secret_key []byte) *service {
	if len(secret_key) < 32 {
		log.Fatal("secret key should more than or same 32 bytes")
	}
	return &service{secret_key}
}

func (s *service) GenerateToken(user users.Users) (string, error) {

	claims := CustomClaims{
		Username: user.Username,
		Email:    user.Email,
		Fullname: user.FullName,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "https://www.blessedenglish.co.id",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(s.secret_key)

	if err != nil {
		return "", errors.New(err.Error())
	}

	return signedToken, nil
}

func (s *service) ValidateToken(token string) (*CustomClaims, error) {

	claims := &CustomClaims{}

	validateToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, err := t.Method.(*jwt.SigningMethodHMAC); !err {
			return nil, errors.New("invalid signin token method")
		}
		return s.secret_key, nil
	})

	if err != nil {
		return nil, errors.New("error parsing token")
	}

	claims, ok := validateToken.Claims.(*CustomClaims)

	if !ok && !validateToken.Valid {

		fmt.Println(ok, validateToken.Valid)
		return nil, errors.New("error token")
	}

	return claims, nil

}

func UsersMidlleware(auth Service, user users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			response := helper.APIResponse(fiber.StatusUnauthorized, "failed", "failed to get authorization header", nil)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse(fiber.StatusUnauthorized, "failed", "failed to get token", err)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		user, err := user.FindUserByUsername(claims.Username)

		if err != nil {
			response := helper.APIResponse(fiber.StatusUnauthorized, "failed", "failed to get user", nil)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		c.Locals("currentUser", user)

		return c.Next()
	}
}

func AdministratorMidlleware(auth Service, user users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		authHeader := c.Get("Authorization")

		if authHeader == "" {
			response := helper.APIResponse(fiber.StatusUnauthorized, "failed", "failed to get authorization header", nil)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		var tokenString string
		_, err := fmt.Sscanf(authHeader, "Bearer %s", &tokenString)

		if err != nil && tokenString == "" {
			response := helper.APIResponse(fiber.StatusUnauthorized, "failed", "failed to get token", nil)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		claims, err := auth.ValidateToken(tokenString)

		if err != nil {
			response := helper.APIResponse(fiber.StatusUnauthorized, "failed", "failed to get token", nil)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		if claims.Role != "administrator" {
			response := helper.APIResponse(fiber.StatusUnauthorized, "failed", "you are not the adaministrator", nil)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		user, err := user.FindUserByUsername(claims.Username)

		if err != nil {
			response := helper.APIResponse(fiber.StatusUnauthorized, "failed", "failed to get user", nil)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		if user.Role != "administrator" {
			response := helper.APIResponse(fiber.StatusUnauthorized, "failed", "user is not administrator", nil)
			return c.Status(fiber.StatusUnauthorized).JSON(response)
		}

		c.Locals("admin", user)

		return c.Next()
	}
}
