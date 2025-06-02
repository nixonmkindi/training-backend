package notification

import (
	"training-backend/server/models"
	"training-backend/services/entity"
)

type Reader interface {
	Get(id int32) (*entity.Notification, error)
	ListByFromID(fromID int32) ([]*entity.Notification, error)
	ListByToID(toID int32) ([]*entity.Notification, error)
	ListByNotificationType(notificationTypeID int32) ([]*entity.Notification, error)
	GetUnreadUserNotification(toID int32) (int32, error)
}

type Writer interface {
	Create(e *entity.Notification) (int32, error)
	UpdateOpenedStatus(id, userID int32) error
	HardDelete(id int32) error
	SoftDelete(id, deletedBy int32) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateOveralNotification(notifications []*models.Notifications) error
	ListNotificationByNotificationType(notificationTypeID int32) ([]*entity.Notification, error)
	GetNotification(id, userID int32) (*entity.Notification, error)
	ListUserNotification(userID int32) (*entity.UserNotification, error)
	SoftDeleteNotification(id, deletedBy int32) error
	HardDeleteNotification(id int32) error
}
