package controllers

import (
	"go-echo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetProduct(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	product := models.Product{}
	product.GetById(id)

	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	name := c.FormValue("name")
	p, _ := strconv.Atoi(c.FormValue("price"))
	price := uint(p)

	product := models.Product{
		Name:  name,
		Price: price,
	}
	product.Create()

	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	name := c.FormValue("name")
	p, _ := strconv.Atoi(c.FormValue("price"))
	price := uint(p)

	product := models.Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
	product.Updates()
	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	product := models.Product{}
	product.DeleteById(id)

	return c.JSON(http.StatusOK, product)
}