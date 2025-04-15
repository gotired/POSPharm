package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gotired/POSPharm/internal/usecase"
)

type Middleware interface {
	JWT(c *fiber.Ctx) error
}

type middlewareProp struct {
	authUsecase usecase.Auth
}

func NewMiddleware(authUsecase usecase.Auth) Middleware {
	return &middlewareProp{authUsecase: authUsecase}
}

func (r *middlewareProp) JWT(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Missing Authorization header"})
	}

	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || strings.ToLower(splitToken[0]) != "bearer" {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token format"})
	}

	decryptToken, err := r.authUsecase.DecryptToken(splitToken[1])
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	claims, err := r.authUsecase.ValidateAccessToken(decryptToken)
	if err != nil || claims == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid token claims"})
	}
	tokenDetail := *claims

	c.Locals("user_id", tokenDetail["user_id"])

	return c.Next()
}
