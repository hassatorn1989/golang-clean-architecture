package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubjectType struct {
	ID        uuid.UUID
	Name      string
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
