package persistence

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/infrastructure/persistence/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubjectTypeGormRepository struct {
	db *gorm.DB
}

func NewSubjectTypeGormRepository(db *gorm.DB) repositories.SubjectTypeRepository {
	return &SubjectTypeGormRepository{db: db}
}

func (r *SubjectTypeGormRepository) FindAll(ctx context.Context) ([]*entities.SubjectType, error) {
	var rows []models.SubjectTypeModel
	if err := r.db.WithContext(ctx).Find(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]*entities.SubjectType, len(rows))
	for i, row := range rows {
		result[i] = mapSubjectTypeModel(row)
	}
	return result, nil
}

func (r *SubjectTypeGormRepository) FindWithPagination(ctx context.Context, limit int, offset int) ([]*entities.SubjectType, error) {
	var rows []models.SubjectTypeModel
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&rows).Error; err != nil {
		return nil, err
	}

	result := make([]*entities.SubjectType, len(rows))
	for i, row := range rows {
		result[i] = mapSubjectTypeModel(row)
	}
	return result, nil
}

func (r *SubjectTypeGormRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.SubjectType, error) {
	var row models.SubjectTypeModel
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&row).Error; err != nil {
		return nil, err
	}
	return mapSubjectTypeModel(row), nil
}

func (r *SubjectTypeGormRepository) Create(ctx context.Context, subjectType *entities.SubjectType) error {
	row := &models.SubjectTypeModel{Name: subjectType.Name}
	return r.db.WithContext(ctx).Create(row).Error
}

func (r *SubjectTypeGormRepository) Update(ctx context.Context, subjectType *entities.SubjectType) error {
	return r.db.WithContext(ctx).Model(&models.SubjectTypeModel{}).Where("id = ?", subjectType.ID).Updates(map[string]interface{}{
		"name": subjectType.Name,
	}).Error
}

func (r *SubjectTypeGormRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.SubjectTypeModel{}, "id = ?", id).Error
}

func mapSubjectTypeModel(row models.SubjectTypeModel) *entities.SubjectType {
	return &entities.SubjectType{
		ID:        row.ID,
		Name:      row.Name,
		CreatedBy: row.CreatedBy,
		UpdatedBy: row.UpdatedBy,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		DeletedAt: row.DeletedAt,
	}
}
