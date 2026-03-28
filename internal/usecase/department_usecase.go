package usecase

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentUseCase struct {
	db             *gorm.DB
	departmentRepo repositories.DepartmentRepository
}

func NewDepartmentUseCase(
	db *gorm.DB,
	departmentRepo repositories.DepartmentRepository,
) *DepartmentUseCase {
	return &DepartmentUseCase{
		db:             db,
		departmentRepo: departmentRepo,
	}
}

func (u *DepartmentUseCase) GetAll(ctx context.Context) ([]*entities.Department, error) {
	departments, err := u.departmentRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return departments, nil
}

func (u *DepartmentUseCase) GetWithPagination(ctx context.Context, limit int, offset int) ([]*entities.Department, error) {
	departments, err := u.departmentRepo.FindWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (u *DepartmentUseCase) GetByID(ctx context.Context, id uuid.UUID) (*entities.Department, error) {
	department, err := u.departmentRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return department, nil
}

func (u *DepartmentUseCase) Create(ctx context.Context, req *dto.DepartmentCreateRequest) error {
	department := &entities.Department{
		Name: req.Name,
	}
	return u.departmentRepo.Create(ctx, department)
}

func (u *DepartmentUseCase) Update(ctx context.Context, id uuid.UUID, req *dto.DepartmentUpdateRequest) error {
	department, err := u.departmentRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	department.Name = req.Name
	return u.departmentRepo.Update(ctx, department)
}

func (u *DepartmentUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.departmentRepo.Delete(ctx, id)
}
