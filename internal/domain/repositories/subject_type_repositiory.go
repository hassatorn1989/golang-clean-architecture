package repositories

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"

	"github.com/google/uuid"
)

type SubjectTypeRepository interface {
	FindAll(ctx context.Context) ([]*entities.SubjectType, error)
	FindWithPagination(ctx context.Context, limit int, offset int) ([]*entities.SubjectType, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.SubjectType, error)
	Create(ctx context.Context, subjectType *entities.SubjectType) error
	Update(ctx context.Context, subjectType *entities.SubjectType) error
	Delete(ctx context.Context, id uuid.UUID) error
}
