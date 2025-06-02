package auth

import (
	"net/http"
	"strings"
	"training-backend/package/log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthJWT() echo.MiddlewareFunc {
	if accessTokenMiddleware != nil {
		return accessTokenMiddleware
	}
	accessTokenMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &JWTCustomClaims{},
		SigningKey:              []byte(GetJWTSecret()),
		TokenLookup:             "header:" + echo.HeaderAuthorization,
		ErrorHandlerWithContext: JWTErrorChecker,
		Skipper:                 SkipperLoginCheck,
	})
	return accessTokenMiddleware
}

// SkipperLoginCheck register all routes that do not need login
func SkipperLoginCheck(c echo.Context) bool {
	if strings.HasSuffix(c.Path(), "/auth") ||
		strings.HasSuffix(c.Path(), "/officers/verify-check-number") {
		return true
	}
	return false
}

// JWTErrorChecker will be executed when user try to access a protected path.
func JWTErrorChecker(err error, c echo.Context) error {
	log.Errorf("error due to jwt checker: %v", err)
	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
