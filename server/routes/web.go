package routes

import (
	"training-backend/package/log"

	"github.com/labstack/echo/v4"
)

// controllersRouters Init Router
func WebRouters(app *echo.Echo) {

	//Protected controllers should be defined in this group
	//This controllers is only accessed by authenticated user
	trainingBackend := app.Group("/training-backend/api/v1") //remove the middleware if you want to make public

	notificationType := trainingBackend.Group("/notification-type")
	{
		log.Info(notificationType)
	}
}
