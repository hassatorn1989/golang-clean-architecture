package entities

import (
	"time"

	"github.com/google/uuid"
)

type Department struct {
	ID        uuid.UUID
	Name      string
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
