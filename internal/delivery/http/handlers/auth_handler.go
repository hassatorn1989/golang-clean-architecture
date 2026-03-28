package handlers

import (
	"golang-clean-architechture/internal/dto"
	"golang-clean-architechture/internal/usecase"
	"golang-clean-architechture/internal/utils"

	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.authUsecase.Register(c.Context(), req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.Success(c, fiber.Map{"message": "register success"})
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	res, err := h.authUsecase.Login(c.Context(), req, c.Get("User-Agent"), c.IP())
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.Success(c, res)
}

func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	var req dto.RefreshRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	res, err := h.authUsecase.RefreshToken(c.Context(), req, c.Get("User-Agent"), c.IP())
	if err != nil {
		return utils.Error(c, fiber.StatusUnauthorized, err.Error())
	}

	return utils.Success(c, res)
}

func (h *AuthHandler) Logout(c fiber.Ctx) error {
	var req dto.LogoutRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "invalid request body")
	}

	if err := h.authUsecase.Logout(c.Context(), req.RefreshToken); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, "logout failed")
	}

	return utils.Success(c, fiber.Map{"message": "logout success"})
}

func (h *AuthHandler) Me(c fiber.Ctx) error {
	return utils.Success(c, fiber.Map{
		"user_id": c.Locals("user_id"),
		"email":   c.Locals("email"),
		"role":    c.Locals("role"),
	})
}
