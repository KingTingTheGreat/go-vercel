package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from JTing")
	})
	e.GET("/api/jting", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello from JTing"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
