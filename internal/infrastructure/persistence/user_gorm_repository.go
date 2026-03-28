package persistence

import (
	"context"
	"golang-clean-architechture/internal/domain/entities"
	"golang-clean-architechture/internal/domain/repositories"
	"golang-clean-architechture/internal/infrastructure/persistence/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) repositories.UserRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	var model models.UserModel
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error; err != nil {
		return nil, err
	}

	return &entities.User{
		ID:        model.ID,
		Email:     model.Email,
		Name:      model.Name,
		Password:  model.Password,
		Role:      model.Role,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

func (r *UserGormRepository) FindAll(ctx context.Context) ([]*entities.User, error) {
	var models []models.UserModel
	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}
	users := make([]*entities.User, len(models))
	for i, model := range models {
		users[i] = &entities.User{
			ID:        model.ID,
			Email:     model.Email,
			Name:      model.Name,
			Password:  model.Password,
			Role:      model.Role,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		}
	}
	return users, nil
}

func (r *UserGormRepository) FindWithPagination(ctx context.Context, limit int, offset int) ([]*entities.User, error) {
	var models []models.UserModel
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&models).Error; err != nil {
		return nil, err
	}
	users := make([]*entities.User, len(models))
	for i, model := range models {
		users[i] = &entities.User{
			ID:        model.ID,
			Email:     model.Email,
			Name:      model.Name,
			Password:  model.Password,
			Role:      model.Role,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		}
	}
	return users, nil
}

func (r *UserGormRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	var model models.UserModel
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}

	return &entities.User{
		ID:        model.ID,
		Email:     model.Email,
		Name:      model.Name,
		Password:  model.Password,
		Role:      model.Role,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

func (r *UserGormRepository) Create(ctx context.Context, user *entities.User) error {
	model := &models.UserModel{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Role:     user.Role,
	}
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *UserGormRepository) Update(ctx context.Context, user *entities.User) error {
	model := &models.UserModel{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Role:     user.Role,
	}
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *UserGormRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.UserModel{}).Error
}
