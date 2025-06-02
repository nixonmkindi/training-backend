package entity

import (
	"errors"
	"time"
	"training-backend/package/log"
)

type Notification struct {
	ID                 int32
	NotificationTypeID int32
	SubsystemID        int32
	FromID             int32
	ToID               int32
	Content            string
	IsOpened           bool
	NotificationType   *NotificationType
	CreatedBy          int32
	CreatedAt          time.Time
	UpdatedBy          int32
	UpdatedAt          time.Time
	DeletedBy          int32
	DeletedAt          time.Time
}

func NewNotification(
	notificationTypeID,
	subsystemID,
	fromID,
	toID int32,
	content string,
	createdBy int32) (*Notification, error) {
	absenteeismStatus := &Notification{
		NotificationTypeID: notificationTypeID,
		SubsystemID:        subsystemID,
		FromID:             fromID,
		ToID:               toID,
		Content:            content,
		CreatedBy:          createdBy,
	}
	err := absenteeismStatus.ValidateNewNotification()
	if err != nil {
		log.Errorf("error validating new Notification entity %v", err)
		return &Notification{}, err
	}

	return absenteeismStatus, err

}

func (r *Notification) ValidateNewNotification() error {
	if r.NotificationTypeID <= 0 {
		return errors.New("error validating Notification entity, notification_type_id field required")
	}
	if r.SubsystemID <= 0 {
		return errors.New("error validating Notification entity, subsystem_id field required")
	}
	if r.FromID <= 0 {
		return errors.New("error validating Notification entity, from_id field required")
	}
	if r.ToID <= 0 {
		return errors.New("error validating Notification entity, to_id field required")
	}
	if r.Content == "" {
		return errors.New("error validating Notification entity, content field required")
	}
	if r.CreatedBy <= 0 {
		return errors.New("error validating Notification entity, created_by field required")
	}
	return nil
}

type UserNotification struct {
	UserID               int32
	TotalUnread          int32
	IncomingNotification []*Notification
	OutgoingNotification []*Notification
}
