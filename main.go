package main

import (
	"github.com/labstack/echo/v4"
	"rabbitmq/routes"
)

func main() {
	e := echo.New()
	routes.Email(e)
	e.Start(":8080")

}
