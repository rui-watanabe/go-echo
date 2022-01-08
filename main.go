package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func main() {

	e := echo.New()

	e.GET("/", hello)

	e.Start(":8080")
}
