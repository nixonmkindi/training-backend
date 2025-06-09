package routes

import (
	"training-backend/package/validator"
	"training-backend/server/middlewares"

	"github.com/labstack/echo/v4"
)

// Routers function
func Routers(app *echo.Echo) {
	//Common middleware for all type of routers
	//app.Use(middlewares.Cors())
	// app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger(true))
	// app.Use(middlewares.Secure())
	// app.Use(middlewares.Recover())
	//app.Use(middlewares.CSRF())
	// app.Use(middlewares.JWT(), middlewares.CheckAuth())
	// app.Use(middlewares.Session())
	// app.Use(auth.KeyAuth())
	app.Validator = validator.GetValidator() //initialize custom validator
	//web routers

	WebRouters(app)
}
