package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func yallo(c echo.Context) error {
	return c.JSON(http.StatusOK, "yallo from the web side!")
}

func test(c echo.Context) error {
	return c.JSON(http.StatusOK, "testing")
}

func main() {
	fmt.Println("Welcome to the server")

	e := echo.New()

	e.GET("/", yallo)
	e.GET("/test", yallo)

	e.Start(":8000")
}
