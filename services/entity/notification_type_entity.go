package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type NotificationType struct {
	ID          int32
	Name        string
	Description string
	CreatedBy   int32
	CreatedAt   time.Time
	UpdatedBy   int32
	UpdatedAt   time.Time
	DeletedBy   int32
	DeletedAt   time.Time
}

func NewNotificationType(name, description string, createdBy int32) (*NotificationType, error) {
	notificationType := &NotificationType{
		Name:        name,
		Description: description,
		CreatedBy:   createdBy,
	}
	err := notificationType.ValidateNewNotificationType()
	if err != nil {
		log.Errorf("error validating new NotificationType entity %v", err)
		return &NotificationType{}, err
	}

	return notificationType, err

}

func (r *NotificationType) ValidateNewNotificationType() error {
	if r.Name == "" {
		return errors.New("error validating NotificationType entity, name field required")
	}
	if r.Description == "" {
		return errors.New("error validating NotificationType entity, description field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating NotificationType entity, created_by field required")
	}
	return nil
}

func (r *NotificationType) ValidateUpdateNotificationType() error {
	if r.ID <= 0 {
		return errors.New("error validating NotificationType entity, id field required")
	}
	if r.Name == "" {
		return errors.New("error validating NotificationType entity, name field required")
	}
	if r.Description == "" {
		return errors.New("error validating NotificationType entity, description field required")
	}
	if r.UpdatedBy <= 0 {
		return errors.New("error validating NotificationType entity, updated_by field required")
	}
	return nil
}
