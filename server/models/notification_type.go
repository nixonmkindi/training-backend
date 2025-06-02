package models

import (
	"time"
)

type NotificationType struct {
	ID          int32     `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	Name        string    `json:"name" form:"name" validate:"required"`
	Description string    `json:"description" form:"description" validate:"required"`
	CreatedBy   int32     `json:"created_by,omitempty" form:"created_by" validate:"numeric"`
	UpdatedBy   int32     `json:"updated_by,omitempty" form:"updated_by" validate:"numeric"`
	DeletedBy   int32     `json:"deleted_by,omitempty" form:"deleted_by" validate:"numeric"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
