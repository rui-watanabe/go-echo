package main

import (
	"go-echo/models"
	"go-echo/routers"
	"net/http"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func main() {
	models.Connect()
	e := routers.NewRouter()
	e.Start(":8080")
}
