package repositories

import (
	"context"

	"golang-clean-architechture/internal/domain/entities"

	"github.com/google/uuid"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindAll(ctx context.Context) ([]*entities.User, error)
	FindWithPagination(ctx context.Context, limit int, offset int) ([]*entities.User, error)
	Create(ctx context.Context, user *entities.User) error
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
