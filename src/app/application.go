package app

import (
	"github.com/labstack/echo/v4"
	"github.com/maei/protocol_buffer_go/src/middleware"
)

var (
	router = echo.New()
)

func StartApplication() {
	mapUrls()
	//router.Use(middleware.HeaderMiddleware.ServerHeader)
	router.Use(middleware.HeaderMiddleware.CheckHeader)
	router.Logger.Fatal(router.Start(":8002"))
}
