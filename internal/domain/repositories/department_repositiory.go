package repositories

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"

	"github.com/google/uuid"
)

type DepartmentRepository interface {
	FindAll(ctx context.Context) ([]*entities.Department, error)
	FindWithPagination(ctx context.Context, limit int, offset int) ([]*entities.Department, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Department, error)
	Create(ctx context.Context, department *entities.Department) error
	Update(ctx context.Context, department *entities.Department) error
	Delete(ctx context.Context, id uuid.UUID) error
}
