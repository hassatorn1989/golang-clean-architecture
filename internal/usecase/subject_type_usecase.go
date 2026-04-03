package usecase

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubjectTypeUseCase struct {
	db              *gorm.DB
	subjectTypeRepo repositories.SubjectTypeRepository
}

func NewSubjectTypeUseCase(db *gorm.DB, subjectTypeRepo repositories.SubjectTypeRepository) *SubjectTypeUseCase {
	return &SubjectTypeUseCase{
		db:              db,
		subjectTypeRepo: subjectTypeRepo,
	}
}

func (u *SubjectTypeUseCase) GetAll(ctx context.Context) ([]*entities.SubjectType, error) {
	return u.subjectTypeRepo.FindAll(ctx)
}

func (u *SubjectTypeUseCase) GetWithPagination(ctx context.Context, limit int, offset int) ([]*entities.SubjectType, error) {
	return u.subjectTypeRepo.FindWithPagination(ctx, limit, offset)
}

func (u *SubjectTypeUseCase) GetByID(ctx context.Context, id uuid.UUID) (*entities.SubjectType, error) {
	return u.subjectTypeRepo.FindByID(ctx, id)
}

func (u *SubjectTypeUseCase) Create(ctx context.Context, req *dto.SubjectTypeCreateRequest) error {
	return u.subjectTypeRepo.Create(ctx, &entities.SubjectType{Name: req.Name})
}

func (u *SubjectTypeUseCase) Update(ctx context.Context, id uuid.UUID, req *dto.SubjectTypeUpdateRequest) error {
	subjectType, err := u.subjectTypeRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}
	subjectType.Name = req.Name
	return u.subjectTypeRepo.Update(ctx, subjectType)
}

func (u *SubjectTypeUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.subjectTypeRepo.Delete(ctx, id)
}
