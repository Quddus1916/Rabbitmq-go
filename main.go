package main

import (
	"github.com/labstack/echo/v4"
	"rabbitmq/config"
	"rabbitmq/routes"
)

func main() {
	e := echo.New()
	routes.Email(e)
	e.Start(":" + config.Getconfig().Port)

}
