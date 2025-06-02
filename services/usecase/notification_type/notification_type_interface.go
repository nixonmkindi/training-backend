package notification_type

import "training-backend/services/entity"

// Reader interface
type Reader interface {
	Get(id int32) (*entity.NotificationType, error)
	List() ([]*entity.NotificationType, error)
	CheckNotificationType(name string) (bool, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.NotificationType) (int32, error)
	Update(e *entity.NotificationType) (int32, error)
	SoftDelete(id, deletedBy int32) error
	HardDelete(id int32) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	CreateNotificationType(name, description string, createdBy int32) (int32, error)
	ListNotificationType() ([]*entity.NotificationType, error)
	GetNotificationType(id int32) (*entity.NotificationType, error)
	UpdateNotificationType(e *entity.NotificationType) (int32, error)
	SoftDeleteNotificationType(id, deletedBy int32) error
	HardDeleteNotificationType(id int32) error
}
