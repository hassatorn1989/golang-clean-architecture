package handlers

import (
	"golang-clean-architechture/internal/dto"
	"golang-clean-architechture/internal/usecase"
	"golang-clean-architechture/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type SubjectTypeHandler struct {
	subjectTypeUseCase *usecase.SubjectTypeUseCase
}

func NewSubjectTypeHandler(subjectTypeUseCase *usecase.SubjectTypeUseCase) *SubjectTypeHandler {
	return &SubjectTypeHandler{subjectTypeUseCase: subjectTypeUseCase}
}

func (h *SubjectTypeHandler) GetAll(c fiber.Ctx) error {
	rows, err := h.subjectTypeUseCase.GetAll(c.Context())
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.Success(c, rows)
}

func (h *SubjectTypeHandler) GetWithPagination(c fiber.Ctx) error {
	limit, offset := utils.GetPaginationParams(c)
	rows, err := h.subjectTypeUseCase.GetWithPagination(c.Context(), limit, offset)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.Success(c, rows)
}

func (h *SubjectTypeHandler) GetByID(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}
	row, err := h.subjectTypeUseCase.GetByID(c.Context(), id)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.Success(c, row)
}

func (h *SubjectTypeHandler) Create(c fiber.Ctx) error {
	var req dto.SubjectTypeCreateRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := h.subjectTypeUseCase.Create(c.Context(), &req); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.Success(c, "Subject type created successfully")
}

func (h *SubjectTypeHandler) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}
	var req dto.SubjectTypeUpdateRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := h.subjectTypeUseCase.Update(c.Context(), id, &req); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.Success(c, "Subject type updated successfully")
}

func (h *SubjectTypeHandler) Delete(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}
	if err := h.subjectTypeUseCase.Delete(c.Context(), id); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.Success(c, "Subject type deleted successfully")
}
