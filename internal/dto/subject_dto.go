package dto

import (
	"time"

	"github.com/google/uuid"
)

type SubjectCreateRequest struct {
	DepartmentID uuid.UUID            `json:"department_id" binding:"required"`
	Title        string               `json:"title" binding:"required"`
	Description  string               `json:"description"`
	Status       string               `json:"status" binding:"required"`
	StartDate    time.Time            `json:"start_date" binding:"required"`
	EndDate      time.Time            `json:"end_date" binding:"required"`
	Year         int                  `json:"year" binding:"required"`
	MaxAnswers   int                  `json:"max_answers" binding:"required"`
	Items        []SubjectItemRequest `json:"items"`
}

type SubjectUpdateRequest struct {
	DepartmentID uuid.UUID            `json:"department_id" binding:"required"`
	Title        string               `json:"title" binding:"required"`
	Description  string               `json:"description"`
	Status       string               `json:"status" binding:"required"`
	StartDate    time.Time            `json:"start_date" binding:"required"`
	EndDate      time.Time            `json:"end_date" binding:"required"`
	Year         int                  `json:"year" binding:"required"`
	MaxAnswers   int                  `json:"max_answers" binding:"required"`
	Items        []SubjectItemRequest `json:"items"`
}

type SubjectItemRequest struct {
	ID          uuid.UUID `json:"id"`
	CategoryID  uuid.UUID `json:"category_id" binding:"required"`
	Description string    `json:"description"`
}
