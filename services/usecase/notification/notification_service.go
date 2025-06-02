package notification

import (
	"training-backend/server/models"
	"training-backend/services/entity"
	"training-backend/services/error_message"
	"training-backend/services/repository"
	"training-backend/services/usecase/email"
	"training-backend/services/usecase/notification_type"
)

type Service struct {
	repo Repository
}

func NewService() UseCase {
	repo := repository.NewNotification()
	return &Service{
		repo: repo,
	}
}
func (s *Service) CreateNotification(
	notificationTypeID,
	subsystemID,
	fromID,
	toID int32,
	userName,
	userEmail,
	content string,
	createdBy int32) (int32, error) {
	var notificationID int32

	notification, err := entity.NewNotification(
		notificationTypeID,
		subsystemID,
		fromID,
		toID,
		content,
		createdBy)
	if err != nil {
		return notificationID, err
	}

	notificationID, err = s.repo.Create(notification)
	if err != nil {
		return notificationID, err
	}

	notificationTypeService := notification_type.NewService()
	notificationType, err := notificationTypeService.GetNotificationType(notificationTypeID)
	if err != nil {
		return notificationID, nil
	}

	email.SendEmailNotification(userName, userEmail, notificationType.Name, content)

	return notificationID, err

}

func (s *Service) CreateOveralNotification(notifications []*models.Notifications) error {

	for _, notification := range notifications {
		for i := 0; i < len(notification.NotificationUser); i++ {
			_, err := s.CreateNotification(
				notification.NotificationTypeID,
				notification.SubsystemID,
				notification.FromID,
				notification.NotificationUser[i].UserID,
				notification.NotificationUser[i].Name,
				notification.NotificationUser[i].Email,
				notification.Content,
				notification.CreatedBy)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *Service) ListNotificationByNotificationType(notificationTypeID int32) ([]*entity.Notification, error) {
	notification, err := s.repo.ListByNotificationType(notificationTypeID)
	if err != nil {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return notification, nil
		}
		return notification, err
	}
	return notification, err
}

func (s *Service) GetNotification(id, userID int32) (*entity.Notification, error) {
	notification, err := s.repo.Get(id)
	if err != nil {
		if err.Error() == error_message.ErrNoResultSet.Error() {
			return notification, nil
		}
		return notification, err
	}

	if notification != nil {
		if notification.ToID == userID {
			err = s.repo.UpdateOpenedStatus(id, userID)
			if err != nil {
				return nil, err
			}
		}
	}
	return notification, err
}

func (s *Service) ListUserNotification(userID int32) (*entity.UserNotification, error) {
	var userNotification *entity.UserNotification

	incomingNotification, err := s.repo.ListByToID(userID)
	if err != nil && err.Error() != error_message.ErrNoResultSet.Error() {
		return userNotification, err
	}

	outgoingNotification, err := s.repo.ListByFromID(userID)
	if err != nil && err.Error() != error_message.ErrNoResultSet.Error() {
		return userNotification, err
	}

	totalUnread, err := s.repo.GetUnreadUserNotification(userID)
	if err != nil && err.Error() != error_message.ErrNoResultSet.Error() {
		return userNotification, err
	}

	userNotification = &entity.UserNotification{
		UserID:               userID,
		TotalUnread:          totalUnread,
		IncomingNotification: incomingNotification,
		OutgoingNotification: outgoingNotification,
	}

	return userNotification, err
}

func (s *Service) SoftDeleteNotification(id, deletedBy int32) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	err = s.repo.SoftDelete(id, deletedBy)
	if err != nil {
		return err
	}
	return err
}

func (s *Service) HardDeleteNotification(id int32) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}
	err = s.repo.HardDelete(id)
	if err != nil {
		return err
	}
	return err
}
