package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rizlantamima/golang-web-boilerplate/config"
)

func init() {
	fmt.Println(`initiate project`)

}

func main() {

	configuration, err := config.New(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	portENV := configuration.GetString("APP_PORT", "3000")
	e.Logger.Fatal(e.Start(":" + portENV))
}
