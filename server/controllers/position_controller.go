package controllers

import (
	"net/http"
	"training-backend/package/log"
	"training-backend/package/trim"
	"training-backend/package/util"
	"training-backend/server/models"
	"training-backend/services/entity"
	"training-backend/services/usecase/position"

	"github.com/labstack/echo/v4"
)

func ListPosition(c echo.Context) error {
	service := position.NewService()
	positions, err := service.ListPosition()
	if util.CheckError(err) {
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	if positions == nil {
		return MessageResponse(c, http.StatusAccepted, "no data found")
	}

	positionResponse := make([]*models.Position, 0)
	for _, d := range positions {
		positionResponse = append(positionResponse, &models.Position{
			ID:          d.ID,
			Name:        d.Name,
			Description: d.Description,
			CreatedBy:   d.CreatedBy,
			CreatedAt:   d.CreatedAt,
			UpdatedBy:   d.UpdatedBy,
			UpdatedAt:   d.UpdatedAt,
			DeletedBy:   d.DeletedBy,
			DeletedAt:   d.DeletedAt,
		})
	}

	return Response(c, http.StatusOK, positionResponse)
}

func ShowPosition(c echo.Context) error {
	service := position.NewService()
	a := &models.Position{}
	if err := c.Bind(&a); util.CheckError(err) {
		log.Errorf("error binding position: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	d, err := service.GetPosition(a.ID)
	if err != nil {
		log.Errorf("error getting getting position: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	positionResponse := models.Position{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		CreatedBy:   d.CreatedBy,
		CreatedAt:   d.CreatedAt,
		UpdatedBy:   d.UpdatedBy,
		UpdatedAt:   d.UpdatedAt,
		DeletedBy:   d.DeletedBy,
		DeletedAt:   d.DeletedAt,
	}

	return Response(c, http.StatusOK, positionResponse)
}

func CreatePosition(c echo.Context) error {
	r := &models.Position{}
	if err := c.Bind(r); util.CheckError(err) {
		log.Errorf("error binding position: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	if err := c.Validate(r); util.CheckError(err) {
		log.Errorf("error validating position: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "error validating position")
	}

	service := position.NewService()
	name := trim.FormatText(r.Name)
	_, err := service.CreatePosition(name, r.Description, r.CreatedBy)
	if util.CheckError(err) {
		log.Errorf("error creating position: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return MessageResponse(c, http.StatusCreated, "position created successfully")
}

func UpdatePosition(c echo.Context) error {
	d := models.Position{}
	if err := c.Bind(&d); util.CheckError(err) {
		log.Errorf("error binding position: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	if err := c.Validate(d); util.CheckError(err) {
		log.Errorf("error validating position:%v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "error validating position")
	}

	service := position.NewService()
	name := trim.FormatText(d.Name)
	data := &entity.Position{
		ID:          d.ID,
		Name:        name,
		Description: d.Description,
		UpdatedBy:   d.UpdatedBy,
	}

	_, err := service.UpdatePosition(data)
	if util.CheckError(err) {
		log.Errorf("error updating position: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return MessageResponse(c, http.StatusAccepted, "position updated successfully")
}

func DeletePosition(c echo.Context) error {
	userService := position.NewService()
	r := &models.Position{}

	if err := c.Bind(&r); util.CheckError(err) {
		log.Errorf("error binding position id: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	err := userService.SoftDeletePosition(r.ID, r.DeletedBy)
	if util.CheckError(err) {
		log.Errorf("error deleting position: %v", err)
		return ErrorResponse(c, http.StatusInternalServerError, "internal server error")
	}

	return MessageResponse(c, http.StatusAccepted, "position deleted successfully")
}
