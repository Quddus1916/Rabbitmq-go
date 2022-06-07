package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SendEmail(c echo.Context) error {
	return c.JSON(http.StatusOK, " sent email")
}
