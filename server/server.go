package server

import (
	"fmt"
	"training-backend/package/config"
	"training-backend/package/log"
	"training-backend/server/routes"
	"training-backend/server/services"

	"github.com/labstack/echo/v4"
)

// StartWebserver starts a webserver
func StartWebserver() {
	// Echo instance
	e := echo.New()
	//Define renderer

	//Disable echo banner
	e.HideBanner = true

	// Routes
	routes.Routers(e)

	//init cache
	services.Init() //check if this solves the problem
	// Start server

	cfg, err := config.New()
	if err != nil {
		log.Errorf("error getting config: %v", err)
	}

	address := fmt.Sprintf("%v:%v", cfg.WebServer.BaseUrl, cfg.WebServer.Port)
	e.Logger.Fatal(e.Start(address))
}
