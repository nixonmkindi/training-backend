package routes

import (
	"training-backend/server/controllers"

	"github.com/labstack/echo/v4"
)

// controllersRouters Init Router
func WebRouters(app *echo.Echo) {

	//Protected controllers should be defined in this group
	//This controllers is only accessed by authenticated user
	notification := app.Group("/training-backend/api/v1") //remove the middleware if you want to make public

	notificationType := notification.Group("/notification-type")
	{
		notificationType.POST("/list", controllers.ListNotificationType)
		notificationType.POST("/create", controllers.CreateNotificationType)
		notificationType.POST("/show", controllers.ShowNotificationType)
		notificationType.POST("/update", controllers.UpdateNotificationType)
		notificationType.POST("/delete", controllers.DeleteNotificationType)
	}

	notifications := notification.Group("/notification")
	{
		notifications.POST("/list", controllers.ListNotificationByNotificationType)
		notifications.POST("/create", controllers.CreateOveralNotification)
		notifications.POST("/show", controllers.ShowNotification)
		notifications.POST("/user-notification", controllers.ListUserNotification)
		notifications.POST("/delete", controllers.SoftDeleteNotification)
		notifications.POST("/hard-delete", controllers.HardDeleteNotification)
	}
}
