package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Position struct {
	ID          int32
	Name        string
	Description string
	CreatedBy   int32
	UpdatedBy   int32
	DeletedBy   int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewPosition(name string, description string, createdBy int32) (*Position, error) {
	entity := &Position{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}

	err := entity.ValidateNewPosition()
	if err != nil {
		log.Errorf("error validating new Position entity %v", err)
		return &Position{}, err
	}

	return entity, nil
}

func (r *Position) ValidateNewPosition() error {
	if r.Name == "" {
		return errors.New("error validating Position entity, name field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Position entity, createdBy field required")
	}
	return nil
}

func (r *Position) ValidateUpdatePosition() error {
	if r.ID <= 0 {
		return errors.New("error validating Position entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating Position entity, name field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating Position entity, updatedBy field required")
	}
	return nil
}
