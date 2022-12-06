package config

import (
	"github.com/labstack/echo/v4"
	"github.com/rizlantamima/golang-web-boilerplate/controllers"
)

func InitRoutes(e *echo.Echo) {

	e.GET("/", controllers.RootPage)

}
