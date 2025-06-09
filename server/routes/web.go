package routes

import (
	"training-backend/server/controllers"

	"github.com/labstack/echo/v4"
)

// controllersRouters Init Router
func WebRouters(app *echo.Echo) {

	//Protected controllers should be defined in this group
	//This controllers is only accessed by authenticated user
	backend := app.Group("/training-backend/api/v1") //remove the middleware if you want to make public

	position := backend.Group("/position")
	{
		position.POST("/list", controllers.ListPosition)
		position.POST("/create", controllers.CreatePosition)
		position.POST("/show", controllers.ShowPosition)
		position.POST("/update", controllers.UpdatePosition)
		position.POST("/delete", controllers.DeletePosition)
	}
}
