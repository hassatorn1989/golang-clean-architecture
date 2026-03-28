package entities

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	TokenHash     string
	FamilyID      uuid.UUID
	ParentTokenID *uuid.UUID
	ExpiresAt     time.Time
	RevokedAt     *time.Time
	ReplacedByID  *uuid.UUID
	UserAgent     string
	IPAddress     string
	CreatedAt     time.Time
}
