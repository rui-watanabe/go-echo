package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Animal struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

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

func addCat(c echo.Context) error {
	cat := Animal{}

	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Failed unmarshaling cat data: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("Cat struct is: %s", cat)
	return c.String(http.StatusOK, "loading cat info is correct")
}
func addDog(c echo.Context) error {
	dog := Animal{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)

	if err != nil {
		log.Printf("Failed unmarshaling dog data: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("Dog struct is: %s", dog)
	return c.String(http.StatusOK, "loading dog info is correct")
}

func addHamsters(c echo.Context) error {
	hamster := Animal{}

	defer c.Request().Body.Close()

	err := c.Bind(&hamster)

	if err != nil {
		log.Printf("Failed unmarshaling hamster data: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("Hamster struct is: %s", hamster)
	return c.String(http.StatusOK, "loading hamster info is correct")
}

func main() {

	e := echo.New()

	e.GET("/", hello)
	e.GET("cats/:data", getCats)
	e.POST("cats", addCat)
	e.POST("dogs", addDog)
	e.POST("hamsters", addHamsters)

	e.Start(":8080")
}
