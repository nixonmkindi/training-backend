package auth

import (
	"training-backend/package/config"
	"training-backend/package/crypto"
	"training-backend/package/log"
	"training-backend/package/util"

	"github.com/labstack/echo/v4"
)

const dataHash = "DATA-HASH"
const dataSignature = "DATA-SIGNATURE"
const systemName = "SYSTEM-NAME"

// KeyAuth middleware
func KeyAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			hash := c.Request().Header.Get(dataHash)
			signature := c.Request().Header.Get(dataSignature)
			system := c.Request().Header.Get(systemName)

			cfg, err := config.New()
			if util.CheckError(err) {
				log.Errorf("error creating config")
				return err
			}
			pubKey, err := cfg.GetSystemPublicKey(system)
			if util.CheckError(err) {
				log.Errorf("error getting public key: %v", err)
				return err
			}
			isValid, err := crypto.Verify(pubKey, hash, signature)
			if isValid && !util.CheckError(err) {
				return next(c)
			}
			return err
		}
	}
}
