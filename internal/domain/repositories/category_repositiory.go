package repositories

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]*entities.Category, error)
	FindWithPagination(ctx context.Context, limit int, offset int) ([]*entities.Category, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Category, error)
	Create(ctx context.Context, category *entities.Category) error
	Update(ctx context.Context, category *entities.Category) error
	Delete(ctx context.Context, id uuid.UUID) error
}
