package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnswerModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	SubjectID uuid.UUID `gorm:"type:uuid;not null;index"`
	AnswerID  uuid.UUID `gorm:"type:uuid;not null;index"`
	CreatedBy uuid.UUID `gorm:"type:uuid;not null"`
	UpdatedBy uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (AnswerModel) TableName() string {
	return "answers"
}

func (d *AnswerModel) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == uuid.Nil {
		d.ID = uuid.New()
	}

	if d.CreatedAt.IsZero() {
		d.CreatedAt = time.Now()
	}

	if d.UpdatedAt.IsZero() {
		d.UpdatedAt = time.Now()
	}

	if d.CreatedBy == uuid.Nil {
		// set data from jwt token
		d.CreatedBy = uuid.New() // placeholder, replace with actual user ID from JWT
		d.UpdatedBy = d.CreatedBy
	}
	return nil
}

func (d *AnswerModel) BeforeUpdate(tx *gorm.DB) (err error) {
	d.UpdatedAt = time.Now()
	// set data from jwt token
	d.UpdatedBy = uuid.New() // placeholder, replace with actual user ID from JWT
	return nil
}

func (d *AnswerModel) BeforeDelete(tx *gorm.DB) (err error) {
	d.DeletedAt.Time = time.Now()
	d.DeletedAt.Valid = true
	return nil
}
