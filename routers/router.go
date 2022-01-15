package routers

import (
	"go-echo/controllers"
	"net/http"

	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/product/:id", controllers.GetProduct)
	e.POST("/api/product", controllers.CreateProduct)
	e.PUT("/api/product/:id", controllers.UpdateProduct)
	e.DELETE("/api/product/:id", controllers.DeleteProduct)

	return e
}
