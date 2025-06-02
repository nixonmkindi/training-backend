package models

import "time"

type Notification struct {
	ID                 int32             `json:"id,omitempty" form:"id" validate:"omitempty,numeric"`
	NotificationTypeID int32             `json:"notification_type_id" form:"notification_type_id"`
	SubsystemID        int32             `json:"subsystem_id" form:"subsystem_id"`
	FromID             int32             `json:"from_id" form:"from_id"`
	ToID               int32             `json:"to_id" form:"to_id"`
	Content            string            `json:"content" form:"content"`
	IsOpened           bool              `json:"is_opened" form:"is_opened"`
	NotificationType   *NotificationType `json:"notification_type"`
	FromUser           string            `json:"from_user"`
	ToUser             string            `json:"to_user"`
	CreatedBy          int32             `json:"created_by,omitempty" form:"created_by" validate:"numeric"`
	UpdaterName        string            `json:"updater_name,omitempty" form:"updater_name"`
	UpdatedBy          int32             `json:"updated_by,omitempty" form:"updated_by" validate:"numeric"`
	DeletedBy          int32             `json:"deleted_by,omitempty" form:"deleted_by" validate:"numeric"`
	CreatedAt          time.Time         `json:"created_at,omitempty"`
	UpdatedAt          time.Time         `json:"updated_at,omitempty"`
	DeletedAt          time.Time         `json:"deleted_at,omitempty"`
}

type Notifications struct {
	NotificationTypeID int32               `json:"notification_type_id" form:"notification_type_id"`
	SubsystemID        int32               `json:"subsystem_id" form:"subsystem_id"`
	FromID             int32               `json:"from_id" form:"from_id"`
	NotificationUser   []*NotificationUser `json:"notification_user"`
	Content            string              `json:"content" form:"content"`
	CreatedBy          int32               `json:"created_by,omitempty" form:"created_by" validate:"numeric"`
}

type UserNotification struct {
	UserID               int32           `json:"user_id"`
	TotalUnread          int32           `json:"total_unread"`
	IncomingNotification []*Notification `json:"incoming_notification"`
	OutgoingNotification []*Notification `json:"outgoing_notification"`
}

type NotificationUser struct {
	UserID int32  `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}
