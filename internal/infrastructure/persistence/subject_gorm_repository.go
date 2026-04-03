package persistence

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/infrastructure/persistence/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubjectGormRepository struct {
	db *gorm.DB
}

func NewSubjectGormRepository(db *gorm.DB) repositories.SubjectRepository {
	return &SubjectGormRepository{db: db}
}

func (r *SubjectGormRepository) FindAll(ctx context.Context) ([]*entities.Subject, error) {
	var models []models.SubjectModel
	if err := r.db.WithContext(ctx).Preload("Items").Find(&models).Error; err != nil {
		return nil, err
	}

	subjects := make([]*entities.Subject, len(models))
	for i, model := range models {
		subjects[i] = mapSubjectModel(model)
	}

	return subjects, nil
}

func (r *SubjectGormRepository) FindWithPagination(ctx context.Context, limit int, offset int) ([]*entities.Subject, error) {
	var models []models.SubjectModel
	if err := r.db.WithContext(ctx).Preload("Items").Limit(limit).Offset(offset).Find(&models).Error; err != nil {
		return nil, err
	}

	subjects := make([]*entities.Subject, len(models))
	for i, model := range models {
		subjects[i] = mapSubjectModel(model)
	}

	return subjects, nil
}

func (r *SubjectGormRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Subject, error) {
	var model models.SubjectModel
	if err := r.db.WithContext(ctx).Preload("Items").Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}

	return mapSubjectModel(model), nil
}

func (r *SubjectGormRepository) Create(ctx context.Context, subject *entities.Subject) error {
	model := &models.SubjectModel{
		DepartmentID: subject.DepartmentID,
		Title:        subject.Title,
		Description:  subject.Description,
		Status:       subject.Status,
		StartDate:    subject.StartDate,
		EndDate:      subject.EndDate,
		Year:         subject.Year,
		MaxAnswers:   subject.MaxAnswers,
		Items:        mapSubjectItemsToModels(subject.Items),
	}
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *SubjectGormRepository) Update(ctx context.Context, subject *entities.Subject) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.SubjectModel{}).Where("id = ?", subject.ID).Updates(map[string]interface{}{
			"department_id": subject.DepartmentID,
			"title":         subject.Title,
			"description":   subject.Description,
			"status":        subject.Status,
			"start_date":    subject.StartDate,
			"end_date":      subject.EndDate,
			"year":          subject.Year,
			"max_answers":   subject.MaxAnswers,
		}).Error; err != nil {
			return err
		}

		if err := tx.Where("subject_id = ?", subject.ID).Delete(&models.SubjectItemModel{}).Error; err != nil {
			return err
		}

		items := mapSubjectItemsToModels(subject.Items)
		if len(items) == 0 {
			return nil
		}

		for i := range items {
			items[i].SubjectID = subject.ID
		}

		return tx.Create(&items).Error
	})
}

func (r *SubjectGormRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.SubjectModel{}, "id = ?", id).Error
}

func mapSubjectModel(model models.SubjectModel) *entities.Subject {
	return &entities.Subject{
		ID:           model.ID,
		DepartmentID: model.DepartmentID,
		Title:        model.Title,
		Description:  model.Description,
		Status:       model.Status,
		StartDate:    model.StartDate,
		EndDate:      model.EndDate,
		Year:         model.Year,
		MaxAnswers:   model.MaxAnswers,
		Items:        mapSubjectItemModels(model.Items),
		CreatedBy:    model.CreatedBy,
		UpdatedBy:    model.UpdatedBy,
		CreatedAt:    model.CreatedAt,
		UpdatedAt:    model.UpdatedAt,
		DeletedAt:    model.DeletedAt,
	}
}

func mapSubjectItemsToModels(items []entities.SubjectItem) []models.SubjectItemModel {
	result := make([]models.SubjectItemModel, len(items))
	for i, item := range items {
		result[i] = models.SubjectItemModel{
			ID:          item.ID,
			CategoryID:  item.CategoryID,
			SubjectID:   item.SubjectID,
			Description: item.Description,
		}
	}
	return result
}

func mapSubjectItemModels(items []models.SubjectItemModel) []entities.SubjectItem {
	result := make([]entities.SubjectItem, len(items))
	for i, item := range items {
		result[i] = entities.SubjectItem{
			ID:          item.ID,
			CategoryID:  item.CategoryID,
			SubjectID:   item.SubjectID,
			Description: item.Description,
			CreatedBy:   item.CreatedBy,
			UpdatedBy:   item.UpdatedBy,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
			DeletedAt:   item.DeletedAt,
		}
	}
	return result
}
