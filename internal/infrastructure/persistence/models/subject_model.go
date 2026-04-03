package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubjectModel struct {
	ID           uuid.UUID          `gorm:"type:uuid;primaryKey"`
	DepartmentID uuid.UUID          `gorm:"type:uuid;not null;index"`
	Title        string             `gorm:"size:255;not null"`
	Description  string             `gorm:"type:text"`
	Status       string             `gorm:"size:50;not null"`
	StartDate    time.Time          `gorm:"not null"`
	EndDate      time.Time          `gorm:"not null"`
	Year         int                `gorm:"not null"`
	MaxAnswers   int                `gorm:"not null"`
	Items        []SubjectItemModel `gorm:"foreignKey:SubjectID"`
	CreatedBy    uuid.UUID          `gorm:"type:uuid;not null"`
	UpdatedBy    uuid.UUID          `gorm:"type:uuid;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (SubjectModel) TableName() string {
	return "subjects"
}

func (s *SubjectModel) BeforeCreate(tx *gorm.DB) (err error) {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}

	if s.CreatedAt.IsZero() {
		s.CreatedAt = time.Now()
	}

	if s.UpdatedAt.IsZero() {
		s.UpdatedAt = time.Now()
	}

	if s.CreatedBy == uuid.Nil {
		// set data from jwt token
		s.CreatedBy = uuid.New() // placeholder, replace with actual user ID from JWT
		s.UpdatedBy = s.CreatedBy
	}
	return nil
}

func (s *SubjectModel) BeforeUpdate(tx *gorm.DB) (err error) {
	s.UpdatedAt = time.Now()
	// set data from jwt token
	s.UpdatedBy = uuid.New() // placeholder, replace with actual user ID from JWT
	return nil
}

func (s *SubjectModel) BeforeDelete(tx *gorm.DB) (err error) {
	s.DeletedAt.Time = time.Now()
	s.DeletedAt.Valid = true
	return nil
}
