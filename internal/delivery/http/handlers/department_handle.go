package handlers

import (
	"golang-clean-architechture/internal/dto"
	"golang-clean-architechture/internal/usecase"
	"golang-clean-architechture/internal/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type DepartmentHandler struct {
	departmentUseCase *usecase.DepartmentUseCase
}

func NewDepartmentHandler(departmentUseCase *usecase.DepartmentUseCase) *DepartmentHandler {
	return &DepartmentHandler{
		departmentUseCase: departmentUseCase,
	}
}

func (h *DepartmentHandler) GetAll(c fiber.Ctx) error {
	departments, err := h.departmentUseCase.GetAll(c.Context())
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, departments)
}

func (h *DepartmentHandler) GetWithPagination(c fiber.Ctx) error {
	limit, offset := utils.GetPaginationParams(c)
	departments, err := h.departmentUseCase.GetWithPagination(c.Context(), limit, offset)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, departments)
}

func (h *DepartmentHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}
	department, err := h.departmentUseCase.GetByID(c.Context(), uuidID)
	if err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, department)
}

func (h *DepartmentHandler) Create(c fiber.Ctx) error {
	var req dto.DepartmentCreateRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := h.departmentUseCase.Create(c.Context(), &req); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}
	return utils.Success(c, "Department created successfully")
}

func (h *DepartmentHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	var req dto.DepartmentUpdateRequest
	if err := c.Bind().Body(&req); err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid request body")
	}
	if err := h.departmentUseCase.Update(c.Context(), uuidID, &req); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, "Department updated successfully")
}

func (h *DepartmentHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return utils.Error(c, fiber.StatusBadRequest, "Invalid ID format")
	}

	if err := h.departmentUseCase.Delete(c.Context(), uuidID); err != nil {
		return utils.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.Success(c, "Department deleted successfully")
}
