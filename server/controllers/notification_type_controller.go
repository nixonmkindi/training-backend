package controllers

import (
	"net/http"
	"training-backend/package/log"
	"training-backend/package/trim"
	"training-backend/package/util"
	"training-backend/server/models"
	"training-backend/services/entity"
	"training-backend/services/usecase/notification_type"

	"github.com/labstack/echo/v4"
)

func ListNotificationType(c echo.Context) error {
	service := notification_type.NewService()
	notificationTypes, err := service.ListNotificationType()
	if util.CheckError(err) {
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}
	if notificationTypes == nil {
		return MessageResponse(c, http.StatusAccepted, "no data found")
	}

	notificationTypeResponse := make([]*models.NotificationType, 0)
	for _, d := range notificationTypes {
		notificationTypeResponse = append(notificationTypeResponse, &models.NotificationType{
			ID:          d.ID,
			Name:        d.Name,
			Description: d.Description,
			CreatedAt:   d.CreatedAt,
			UpdatedBy:   d.UpdatedBy,
			UpdatedAt:   d.UpdatedAt,
			DeletedBy:   d.DeletedBy,
			DeletedAt:   d.DeletedAt,
		})
	}
	return Response(c, http.StatusOK, notificationTypeResponse)
}

func ShowNotificationType(c echo.Context) error {
	service := notification_type.NewService()
	a := &models.NotificationType{}
	if err := c.Bind(&a); util.CheckError(err) {
		log.Errorf("error binding notification type:%v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}

	d, err := service.GetNotificationType(a.ID)
	if err != nil {
		log.Errorf("error getting getting notification type:%v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}

	notificationTypeResponse := models.NotificationType{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedBy:   d.UpdatedBy,
		UpdatedAt:   d.UpdatedAt,
		DeletedBy:   d.DeletedBy,
		DeletedAt:   d.DeletedAt,
	}
	return Response(c, http.StatusOK, notificationTypeResponse)
}
func CreateNotificationType(c echo.Context) error {
	r := &models.NotificationType{}
	if err := c.Bind(r); util.CheckError(err) {
		log.Errorf("error binding notification type: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}
	if err := c.Validate(r); util.CheckError(err) {
		log.Errorf("error validating notification type: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorValidation")
	}
	service := notification_type.NewService()

	name := trim.FormatText(r.Name)
	_, err := service.CreateNotificationType(name, r.Description, r.CreatedBy)
	if util.CheckError(err) {
		log.Errorf("error occurred:%v\n", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}
	return MessageResponse(c, http.StatusCreated, "notification type created successfully")
}

func UpdateNotificationType(c echo.Context) error {
	d := models.NotificationType{}
	if err := c.Bind(&d); util.CheckError(err) {
		log.Errorf("error binding notification type:%v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}
	if err := c.Validate(d); util.CheckError(err) {
		log.Errorf("error validating notification type:%v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorValidation")
	}
	service := notification_type.NewService()

	name := trim.FormatText(d.Name)
	data := &entity.NotificationType{
		ID:          d.ID,
		Name:        name,
		Description: d.Description,
		UpdatedBy:   d.UpdatedBy,
	}
	_, err := service.UpdateNotificationType(data)
	if util.CheckError(err) {
		log.Errorf("error updating notification type: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}
	return MessageResponse(c, http.StatusAccepted, "notification type updated successfully")
}

func DeleteNotificationType(c echo.Context) error {
	userService := notification_type.NewService()
	r := &models.NotificationType{}

	if err := c.Bind(&r); util.CheckError(err) {
		log.Errorf("error binding notification type id: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}

	err := userService.SoftDeleteNotificationType(r.ID, r.DeletedBy)
	if util.CheckError(err) {
		log.Errorf("error deleting notification type:%v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "ErrorInternal")
	}
	return MessageResponse(c, http.StatusAccepted, "notification type deleted sucessfuly")
}
