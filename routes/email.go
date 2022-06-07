package routes

import (
	"github.com/labstack/echo/v4"
	"rabbitmq/controller"
)

func Email(e *echo.Echo) {
	e.GET("/send-email", controller.SendEmail)

}
