package main

import (
	"fmt"
	"log"

	"github.com/anggunpermata/integration-webhook/config"
	"github.com/anggunpermata/integration-webhook/controller"
	"github.com/labstack/echo"
)

func main() {
	log.Println("Starting the HTTP server on port 8080")
	e := echo.New()
	config.InitPort()
	Routes(e)
	Port := fmt.Sprintf(":%d", config.PORT)
	fmt.Println("Port : ", Port)
	if err := e.Start(Port); err != nil {
		e.Logger.Fatal(err)
	}
}

func Routes(e *echo.Echo) {
	e.POST("/caa", controller.AssignAgentWebhook)
}
