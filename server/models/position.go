package models

import (
	"time"
)

type Position struct {
	ID          int32     `json:"id" validate:"omitempty,numeric"`
	Name        string    `json:"name,omitempty" validate:"required"`
	Description string    `json:"description,omitempty" validate:"omitempty"`
	CreatedBy   int32     `json:"created_by,omitempty" validate:"omitempty,numeric"`
	UpdatedBy   int32     `json:"updated_by,omitempty" validate:"omitempty,numeric"`
	DeletedBy   int32     `json:"deleted_by,omitempty" validate:"omitempty,numeric"`
	CreatedAt   time.Time `json:"created_at,omitempty" validate:"omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" validate:"omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty" validate:"omitempty"`
}
