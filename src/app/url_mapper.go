package app

import "github.com/maei/protocol_buffer_go/src/controller"

func publicRoutes() {
	router.GET("/ping", controller.PingController.Ping)
	router.POST("/write", controller.AddressController.WriteAddress)
	router.GET("/write", controller.AddressController.ReadAddress)

}

func mapUrls() {
	publicRoutes()
}
