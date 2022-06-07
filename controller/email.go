package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"rabbitmq/helper"
)

func SendEmail(c echo.Context) error {
	email := c.Param("email")
	res := helper.SendEmail(email)
	return c.JSON(http.StatusOK, res)
}
