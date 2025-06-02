package notification_type

import (
	"training-backend/package/log"
	"training-backend/services/entity"
	"training-backend/services/error_message"
	"training-backend/services/repository"
)

// Service Initialize repository
type Service struct {
	repo Repository
}

// NewService Instantiate new service
func NewService() UseCase {
	repo := repository.NewNotificationType()
	return &Service{
		repo: repo,
	}
}

// CreateNotificationType Calls create new record repository
func (s *Service) CreateNotificationType(name, description string, createdBy int32) (int32, error) {

	var notificationTypeID int32

	notificationType, err := entity.NewNotificationType(name, description, createdBy)
	if err != nil {
		log.Error(err)
		return notificationTypeID, err
	}

	exists, _ := s.repo.CheckNotificationType(name)
	if !exists {
		notificationTypeID, err = s.repo.Create(notificationType)
		if err != nil {
			log.Errorf("error creating notification type %v", err)
			return notificationTypeID, error_message.ErrCannotBeCreated
		}
	}

	return notificationTypeID, err
}

// ListNotificationType Calls list records repository
func (s *Service) ListNotificationType() ([]*entity.NotificationType, error) {
	notificationType, err := s.repo.List()
	if err != nil && err.Error() != error_message.ErrNoResultSet.Error() {
		log.Error(err)
		return notificationType, err
	}
	return notificationType, err
}

// GetNotificationType Calls get single record repository
func (s *Service) GetNotificationType(id int32) (*entity.NotificationType, error) {
	notificationType, err := s.repo.Get(id)
	if err != nil && err.Error() != error_message.ErrNoResultSet.Error() {
		log.Error(err)
		return notificationType, err
	}
	return notificationType, err
}

// UpdateNotificationType Calls updates single record by ID field repository
func (s *Service) UpdateNotificationType(e *entity.NotificationType) (int32, error) {
	err := e.ValidateUpdateNotificationType()
	if err != nil {
		log.Error(err)
		return e.ID, error_message.ErrCannotBeUpdated
	}
	_, err = s.repo.Update(e)
	if err != nil {
		log.Error(err)
		return e.ID, error_message.ErrNotFound
	}
	return e.ID, err
}

// SoftDeleteNotificationType Calls soft delete function for single record by ID repository
func (s *Service) SoftDeleteNotificationType(id, deletedBy int32) error {
	_, err := s.GetNotificationType(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	errDelete := s.repo.SoftDelete(id, deletedBy)
	if errDelete != nil {
		log.Error(errDelete)
		return error_message.ErrCannotBeDeleted
	}
	return errDelete
}

// HardDeleteNotificationType Calls hard delete function for single record by ID repository
func (s *Service) HardDeleteNotificationType(id int32) error {
	_, err := s.GetNotificationType(id)
	if err != nil {
		log.Error(err)
		return error_message.ErrCannotBeDeleted
	}
	errDelete := s.repo.HardDelete(id)
	if errDelete != nil {
		log.Error(errDelete)
		return error_message.ErrCannotBeDeleted
	}
	return errDelete
}
