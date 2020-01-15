package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
    "github.com/sirupsen/logrus" 
)

func yallo(c echo.Context) error {
	return c.JSON(http.StatusOK, "yallo from the web side!")
}

func test(c echo.Context) error {
	return c.JSON(http.StatusOK, "testing")
}

func main() {
	var log = logrus.New()
	fmt.Println("Welcome to the server")
	log.Info("Details")

	e := echo.New()

	e.GET("/", yallo)
	e.GET("/test", test)

	e.Start(":8000")
}
