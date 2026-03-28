package handlers

import (
	"golang-clean-architechture/internal/dto"
	"golang-clean-architechture/internal/usecase"
	"golang-clean-architechture/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type CategoryHandler struct {
	categoryUseCase *usecase.CategoryUseCase
}

func NewCategoryHandler(categoryUseCase *usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{
		categoryUseCase: categoryUseCase,
	}
}

func (h *CategoryHandler) GetAll(c fiber.Ctx) error {
	categorys, err := h.categoryUseCase.GetAll(c.Context())
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, categorys)
}

func (h *CategoryHandler) GetWithPagination(c fiber.Ctx) error {
	limit, offset := utils.GetPaginationParams(c)
	categorys, err := h.categoryUseCase.GetWithPagination(c.Context(), limit, offset)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, categorys)
}

func (h *CategoryHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}
	category, err := h.categoryUseCase.GetByID(c.Context(), uuidID)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, category)
}

func (h *CategoryHandler) Create(c fiber.Ctx) error {
	var req dto.CategoryCreateRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := h.categoryUseCase.Create(c.Context(), &req); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.Success(c, "Category created successfully")
}

func (h *CategoryHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	var req dto.CategoryUpdateRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := h.categoryUseCase.Update(c.Context(), uuidID, &req); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, "Category updated successfully")
}

func (h *CategoryHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	if err := h.categoryUseCase.Delete(c.Context(), uuidID); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, "Category deleted successfully")
}
