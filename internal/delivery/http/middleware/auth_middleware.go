package middleware

import (
	"golang-clean-architechture/internal/domain/services"
	"strings"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(tokenSvc services.TokenService) fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "missing authorization header"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid authorization header"})
		}

		payload, err := tokenSvc.ParseAccessToken(c.Context(), parts[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "invalid access token"})
		}

		c.Locals("user_id", payload.UserID)
		c.Locals("email", payload.Email)
		c.Locals("role", payload.Role)

		return c.Next()
	}
}
