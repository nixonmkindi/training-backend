package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Cors Middleware
func Cors() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		//AllowOrigins: []string{"http://172.16.1.166:3000", "*"},
		AllowOrigins: []string{"http://localhost:3000", "*"},
		AllowMethods: []string{echo.HEAD, echo.GET, echo.PUT, echo.POST, echo.DELETE},
	})
}
