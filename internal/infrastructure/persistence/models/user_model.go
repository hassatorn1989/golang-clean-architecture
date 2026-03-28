package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey"`
	Email     string    `gorm:"size:150;uniqueIndex;not null"`
	Name      string    `gorm:"size:150;not null"`
	Password  string    `gorm:"size:255;not null"`
	Role      string    `gorm:"size:50;not null;default:user"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}

	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}

	if u.UpdatedAt.IsZero() {
		u.UpdatedAt = time.Now()
	}
	return nil
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = time.Now()
	return nil
}

func (UserModel) TableName() string {
	return "users"
}
