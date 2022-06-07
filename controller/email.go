package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"rabbitmq/helper"
)

func SendEmail(c echo.Context) error {
	email := c.Param("email")
	res := helper.PublishEmail(email)
	fmt.Printf("sent email to %v", email)
	return c.JSON(http.StatusOK, res)
}
