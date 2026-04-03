package usecase

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubjectUseCase struct {
	db          *gorm.DB
	subjectRepo repositories.SubjectRepository
}

func NewSubjectUseCase(
	db *gorm.DB,
	subjectRepo repositories.SubjectRepository,
) *SubjectUseCase {
	return &SubjectUseCase{
		db:          db,
		subjectRepo: subjectRepo,
	}
}

func (u *SubjectUseCase) GetAll(ctx context.Context) ([]*entities.Subject, error) {
	subjects, err := u.subjectRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return subjects, nil
}

func (u *SubjectUseCase) GetWithPagination(ctx context.Context, limit int, offset int) ([]*entities.Subject, error) {
	subjects, err := u.subjectRepo.FindWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	return subjects, nil
}

func (u *SubjectUseCase) GetByID(ctx context.Context, id uuid.UUID) (*entities.Subject, error) {
	subject, err := u.subjectRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return subject, nil
}

func (u *SubjectUseCase) Create(ctx context.Context, req *dto.SubjectCreateRequest) error {
	subject := &entities.Subject{
		DepartmentID: req.DepartmentID,
		Title:        req.Title,
		Description:  req.Description,
		Status:       req.Status,
		StartDate:    req.StartDate,
		EndDate:      req.EndDate,
		Year:         req.Year,
		MaxAnswers:   req.MaxAnswers,
		Items:        mapSubjectItemRequests(req.Items),
	}
	return u.subjectRepo.Create(ctx, subject)
}

func (u *SubjectUseCase) Update(ctx context.Context, id uuid.UUID, req *dto.SubjectUpdateRequest) error {
	subject, err := u.subjectRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// delete existing subject items
	for _, item := range subject.Items {
		if err := u.subjectRepo.Delete(ctx, item.ID); err != nil {
			return err
		}
	}

	subject.DepartmentID = req.DepartmentID
	subject.Title = req.Title
	subject.Description = req.Description
	subject.Status = req.Status
	subject.StartDate = req.StartDate
	subject.EndDate = req.EndDate
	subject.Year = req.Year
	subject.MaxAnswers = req.MaxAnswers
	subject.Items = mapSubjectItemRequests(req.Items)
	return u.subjectRepo.Update(ctx, subject)
}

func (u *SubjectUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.subjectRepo.Delete(ctx, id)
}

func mapSubjectItemRequests(items []dto.SubjectItemRequest) []entities.SubjectItem {
	result := make([]entities.SubjectItem, len(items))
	for i, item := range items {
		result[i] = entities.SubjectItem{
			ID:          item.ID,
			CategoryID:  item.CategoryID,
			Description: item.Description,
		}
	}
	return result
}
