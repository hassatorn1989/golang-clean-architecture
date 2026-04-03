package handlers

import (
	"golang-clean-architechture/internal/dto"
	"golang-clean-architechture/internal/usecase"
	"golang-clean-architechture/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type SubjectHandler struct {
	subjectUseCase *usecase.SubjectUseCase
}

func NewSubjectHandler(subjectUseCase *usecase.SubjectUseCase) *SubjectHandler {
	return &SubjectHandler{
		subjectUseCase: subjectUseCase,
	}
}

func (h *SubjectHandler) GetAll(c fiber.Ctx) error {
	subjects, err := h.subjectUseCase.GetAll(c.Context())
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, subjects)
}

func (h *SubjectHandler) GetWithPagination(c fiber.Ctx) error {
	limit, offset := utils.GetPaginationParams(c)
	subjects, err := h.subjectUseCase.GetWithPagination(c.Context(), limit, offset)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, subjects)
}

func (h *SubjectHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	subject, err := h.subjectUseCase.GetByID(c.Context(), uuidID)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, subject)
}

func (h *SubjectHandler) Create(c fiber.Ctx) error {
	var req dto.SubjectCreateRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := h.subjectUseCase.Create(c.Context(), &req); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, "Subject created successfully")
}

func (h *SubjectHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	var req dto.SubjectUpdateRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := h.subjectUseCase.Update(c.Context(), uuidID, &req); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, "Subject updated successfully")
}

func (h *SubjectHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	if err := h.subjectUseCase.Delete(c.Context(), uuidID); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, "Subject deleted successfully")
}
