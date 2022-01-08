package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")
}

func getCats(c echo.Context) error {
	n := c.QueryParam("name")
	t := c.QueryParam("type")

	d := c.Param("data")

	if d == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Cat name is %s\n and type is %s\n", n, t))
	}

	if d == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": n,
			"type": t,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "you need to lets us know if you want json or string data",
	})
}

func main() {

	e := echo.New()

	e.GET("/", hello)
	e.GET("cats/:data", getCats)

	e.Start(":8080")
}
