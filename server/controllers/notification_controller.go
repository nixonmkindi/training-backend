package controllers

import (
	"net/http"
	"training-backend/package/log"
	"training-backend/package/util"
	"training-backend/server/models"
	"training-backend/services/usecase/notification"

	"github.com/labstack/echo/v4"
)

func ListNotificationByNotificationType(c echo.Context) error {
	service := notification.NewService()
	a := &models.ID{}
	if err := c.Bind(&a); util.CheckError(err) {
		log.Errorf("error binding notification type:%v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}

	if err := c.Validate(a); util.CheckError(err) {
		log.Errorf("error validating notification type: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorValidation")
	}

	notifications, err := service.ListNotificationByNotificationType(a.ID)
	if util.CheckError(err) {
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}
	if notifications == nil {
		return MessageResponse(c, http.StatusAccepted, "no data found")
	}

	notificationResponse := make([]*models.Notification, 0)
	for _, notification := range notifications {

		notificationType := &models.NotificationType{
			ID:   notification.NotificationType.ID,
			Name: notification.NotificationType.Name,
		}

		notificationResponse = append(notificationResponse, &models.Notification{
			ID:                 notification.ID,
			NotificationTypeID: notification.NotificationTypeID,
			SubsystemID:        notification.SubsystemID,
			FromID:             notification.FromID,
			ToID:               notification.ToID,
			Content:            notification.Content,
			IsOpened:           notification.IsOpened,
			NotificationType:   notificationType,
			CreatedAt:          notification.CreatedAt,
			UpdatedAt:          notification.UpdatedAt,
		})
	}
	return Response(c, http.StatusOK, notificationResponse)
}

func ShowNotification(c echo.Context) error {
	service := notification.NewService()
	a := &models.Notification{}
	if err := c.Bind(&a); util.CheckError(err) {
		log.Errorf("error binding notification type:%v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}

	notification, err := service.GetNotification(a.ID, a.ToID)
	if err != nil {
		log.Errorf("error getting getting notification type:%v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}

	notificationType := &models.NotificationType{
		ID:   notification.NotificationType.ID,
		Name: notification.NotificationType.Name,
	}

	notificationResponse := models.Notification{
		ID:                 notification.ID,
		NotificationTypeID: notification.NotificationTypeID,
		SubsystemID:        notification.SubsystemID,
		FromID:             notification.FromID,
		ToID:               notification.ToID,
		Content:            notification.Content,
		IsOpened:           notification.IsOpened,
		NotificationType:   notificationType,
		CreatedAt:          notification.CreatedAt,
		UpdatedAt:          notification.UpdatedAt,
	}
	return Response(c, http.StatusOK, notificationResponse)
}

func CreateOveralNotification(c echo.Context) error {
	notificationModel := []*models.Notifications{}
	if err := c.Bind(&notificationModel); util.CheckError(err) {
		log.Errorf("error binding notification fields: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	service := notification.NewService()
	err := service.CreateOveralNotification(notificationModel)
	if util.CheckError(err) {
		log.Errorf("error creating new notification: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return MessageResponse(c, http.StatusOK, "notification created successfully")
}

func ListUserNotification(c echo.Context) error {
	modelID := &models.ID{}
	if err := c.Bind(&modelID); util.CheckError(err) {
		log.Errorf("error binding notification id: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	if err := c.Validate(modelID); util.CheckError(err) {
		log.Errorf("error validating notification id: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "error validating notification id")
	}

	service := notification.NewService()
	notifications, err := service.ListUserNotification(modelID.ID)
	if util.CheckError(err) {
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
	if notifications == nil {
		return MessageResponse(c, http.StatusOK, "no notifications  data found")
	}

	incomingNotifications := make([]*models.Notification, 0)
	for _, notification := range notifications.IncomingNotification {

		notificationType := &models.NotificationType{
			ID:   notification.NotificationType.ID,
			Name: notification.NotificationType.Name,
		}

		incomingNotifications = append(incomingNotifications, &models.Notification{
			ID:                 notification.ID,
			NotificationTypeID: notification.NotificationTypeID,
			SubsystemID:        notification.SubsystemID,
			FromID:             notification.FromID,
			ToID:               notification.ToID,
			Content:            notification.Content,
			IsOpened:           notification.IsOpened,
			NotificationType:   notificationType,
			CreatedAt:          notification.CreatedAt,
			UpdatedAt:          notification.UpdatedAt,
		})
	}

	outgoingNotification := make([]*models.Notification, 0)
	for _, notification := range notifications.OutgoingNotification {

		notificationType := &models.NotificationType{
			ID:   notification.NotificationType.ID,
			Name: notification.NotificationType.Name,
		}

		outgoingNotification = append(outgoingNotification, &models.Notification{
			ID:                 notification.ID,
			NotificationTypeID: notification.NotificationTypeID,
			SubsystemID:        notification.SubsystemID,
			FromID:             notification.FromID,
			ToID:               notification.ToID,
			Content:            notification.Content,
			IsOpened:           notification.IsOpened,
			NotificationType:   notificationType,
			CreatedAt:          notification.CreatedAt,
			UpdatedAt:          notification.UpdatedAt,
		})
	}

	userNotification := &models.UserNotification{
		UserID:               modelID.ID,
		TotalUnread:          notifications.TotalUnread,
		IncomingNotification: incomingNotifications,
		OutgoingNotification: outgoingNotification,
	}
	return Response(c, http.StatusOK, userNotification)
}

func SoftDeleteNotification(c echo.Context) error {

	notificationDeletedBy := &models.DeletedBy{}

	if err := c.Bind(&notificationDeletedBy); util.CheckError(err) {
		log.Errorf("error binding notification fields: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
	if err := c.Validate(notificationDeletedBy); util.CheckError(err) {
		log.Errorf("error validating notification fields: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "error validating notification")
	}

	service := notification.NewService()
	err := service.SoftDeleteNotification(notificationDeletedBy.ID, notificationDeletedBy.DeletedBy)
	if util.CheckError(err) {
		log.Errorf("error deleting notification: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return MessageResponse(c, http.StatusOK, "notification deleted successfully")
}

func HardDeleteNotification(c echo.Context) error {

	modelID := &models.ID{}
	if err := c.Bind(&modelID); util.CheckError(err) {
		log.Errorf("error binding notification id: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
	if err := c.Validate(modelID); util.CheckError(err) {
		log.Errorf("error validating notification id: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "error validating notification id")
	}

	service := notification.NewService()
	err := service.HardDeleteNotification(modelID.ID)
	if util.CheckError(err) {
		log.Errorf("error deleting notification: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}
	return MessageResponse(c, http.StatusOK, "notification deleted successfully")
}
