package persistence

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/infrastructure/persistence/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentGormRepository struct {
	db *gorm.DB
}

func NewDepartmentGormRepository(db *gorm.DB) repositories.DepartmentRepository {
	return &DepartmentGormRepository{db: db}
}

func (r *DepartmentGormRepository) FindAll(ctx context.Context) ([]*entities.Department, error) {
	var models []models.DepartmentModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	departments := make([]*entities.Department, len(models))
	for i, model := range models {
		departments[i] = &entities.Department{
			ID:        model.ID,
			Name:      model.Name,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		}
	}
	return departments, nil

}

func (r *DepartmentGormRepository) FindWithPagination(ctx context.Context, limit int, offset int) ([]*entities.Department, error) {
	var models []models.DepartmentModel
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&models).Error; err != nil {
		return nil, err
	}

	departments := make([]*entities.Department, len(models))
	for i, model := range models {
		departments[i] = &entities.Department{
			ID:        model.ID,
			Name:      model.Name,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		}
	}
	return departments, nil
}

func (r *DepartmentGormRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Department, error) {
	var model models.DepartmentModel
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &entities.Department{
		ID:        model.ID,
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

func (r *DepartmentGormRepository) Create(ctx context.Context, department *entities.Department) error {
	model := &models.DepartmentModel{
		Name: department.Name,
	}
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *DepartmentGormRepository) Update(ctx context.Context, department *entities.Department) error {
	return r.db.WithContext(ctx).Model(&models.DepartmentModel{}).Where("id = ?", department.ID).Updates(map[string]interface{}{
		"name": department.Name,
	}).Error
}

func (r *DepartmentGormRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.DepartmentModel{}, "id = ?", id).Error
}
