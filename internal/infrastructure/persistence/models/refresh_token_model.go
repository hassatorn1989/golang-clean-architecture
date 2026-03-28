package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RefreshTokenModel struct {
	ID            uuid.UUID  `gorm:"type:char(36);primaryKey"`
	UserID        uuid.UUID  `gorm:"type:char(36);not null;index"`
	TokenHash     string     `gorm:"size:255;not null;uniqueIndex"`
	FamilyID      uuid.UUID  `gorm:"type:char(36);not null;index"`
	ParentTokenID *uuid.UUID `gorm:"type:char(36);index"`
	ExpiresAt     time.Time  `gorm:"not null;index"`
	RevokedAt     *time.Time `gorm:"index"`
	ReplacedByID  *uuid.UUID `gorm:"type:char(36)"`
	UserAgent     string     `gorm:"size:255"`
	IPAddress     string     `gorm:"size:100"`
	CreatedAt     time.Time

	User UserModel `gorm:"foreignKey:UserID;references:ID"`
}

func (r *RefreshTokenModel) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == uuid.Nil {
		r.ID = uuid.New()
	}
	return nil
}

func (RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}
