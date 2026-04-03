package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subject struct {
	ID           uuid.UUID
	DepartmentID uuid.UUID
	Title        string
	Description  string
	Status       string
	StartDate    time.Time
	EndDate      time.Time
	Year         int
	MaxAnswers   int
	Items        []SubjectItem
	CreatedBy    uuid.UUID
	UpdatedBy    uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type SubjectItem struct {
	ID          uuid.UUID
	CategoryID  uuid.UUID
	SubjectID   uuid.UUID
	Description string
	CreatedBy   uuid.UUID
	UpdatedBy   uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
