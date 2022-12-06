package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RootPage(c echo.Context) error {
	return c.String(http.StatusOK, "This Is Your Root Page!")
}
