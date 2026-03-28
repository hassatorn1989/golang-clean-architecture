package persistence

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/infrastructure/persistence/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryGormRepository struct {
	db *gorm.DB
}

func NewCategoryGormRepository(db *gorm.DB) repositories.CategoryRepository {
	return &CategoryGormRepository{db: db}
}

func (r *CategoryGormRepository) FindAll(ctx context.Context) ([]*entities.Category, error) {
	var models []models.CategoryModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	categorys := make([]*entities.Category, len(models))
	for i, model := range models {
		categorys[i] = &entities.Category{
			ID:        model.ID,
			Code:      model.Code,
			Name:      model.Name,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		}
	}
	return categorys, nil

}

func (r *CategoryGormRepository) FindWithPagination(ctx context.Context, limit int, offset int) ([]*entities.Category, error) {
	var models []models.CategoryModel
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&models).Error; err != nil {
		return nil, err
	}

	categorys := make([]*entities.Category, len(models))
	for i, model := range models {
		categorys[i] = &entities.Category{
			ID:        model.ID,
			Code:      model.Code,
			Name:      model.Name,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		}
	}
	return categorys, nil
}

func (r *CategoryGormRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.Category, error) {
	var model models.CategoryModel
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	return &entities.Category{
		ID:        model.ID,
		Code:      model.Code,
		Name:      model.Name,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

func (r *CategoryGormRepository) Create(ctx context.Context, category *entities.Category) error {
	model := &models.CategoryModel{
		Code: category.Code,
		Name: category.Name,
	}
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *CategoryGormRepository) Update(ctx context.Context, category *entities.Category) error {
	return r.db.WithContext(ctx).Model(&models.CategoryModel{}).Where("id = ?", category.ID).Updates(map[string]interface{}{
		"code": category.Code,
		"name": category.Name,
	}).Error
}

func (r *CategoryGormRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.CategoryModel{}, "id = ?", id).Error
}
