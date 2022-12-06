package main

import (
	"fmt"
	"log"

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

	config.InitRoutes(e)

	portENV := configuration.GetString("APP_PORT", "3000")
	e.Logger.Fatal(e.Start(":" + portENV))
}
