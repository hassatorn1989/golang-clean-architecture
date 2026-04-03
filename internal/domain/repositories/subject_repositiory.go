package repositories

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"

	"github.com/google/uuid"
)

type SubjectRepository interface {
	FindAll(ctx context.Context) ([]*entities.Subject, error)
	FindWithPagination(ctx context.Context, limit int, offset int) ([]*entities.Subject, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.Subject, error)
	Create(ctx context.Context, subject *entities.Subject) error
	Update(ctx context.Context, subject *entities.Subject) error
	Delete(ctx context.Context, id uuid.UUID) error
}
