package usecase

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryUseCase struct {
	db           *gorm.DB
	categoryRepo repositories.CategoryRepository
}

func NewCategoryUseCase(
	db *gorm.DB,
	categoryRepo repositories.CategoryRepository,
) *CategoryUseCase {
	return &CategoryUseCase{
		db:           db,
		categoryRepo: categoryRepo,
	}
}

func (u *CategoryUseCase) GetAll(ctx context.Context) ([]*entities.Category, error) {
	categorys, err := u.categoryRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (u *CategoryUseCase) GetWithPagination(ctx context.Context, limit int, offset int) ([]*entities.Category, error) {
	categorys, err := u.categoryRepo.FindWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return categorys, nil
}

func (u *CategoryUseCase) GetByID(ctx context.Context, id uuid.UUID) (*entities.Category, error) {
	category, err := u.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (u *CategoryUseCase) Create(ctx context.Context, req *dto.CategoryCreateRequest) error {
	category := &entities.Category{
		Name: req.Name,
	}
	return u.categoryRepo.Create(ctx, category)
}

func (u *CategoryUseCase) Update(ctx context.Context, id uuid.UUID, req *dto.CategoryUpdateRequest) error {
	category, err := u.categoryRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	category.Name = req.Name
	return u.categoryRepo.Update(ctx, category)
}

func (u *CategoryUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.categoryRepo.Delete(ctx, id)
}
