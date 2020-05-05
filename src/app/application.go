package app

import (
	"github.com/labstack/echo/v4"
	serverheader "github.com/maei/protocol_buffer_go/src/middleware"
)

var (
	router = echo.New()
)

func StartApplication() {
	mapUrls()
	router.Use(serverheader.ServerHeader)
	router.Logger.Fatal(router.Start(":8002"))
}
