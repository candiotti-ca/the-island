package main

import (
	"net/http"
	"theisland"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/map", generateMap)

	e.Logger.Fatal(e.Start(":8080"))
}

func generateMap(c echo.Context) error {
	var seed int64 = 5
	generation := theisland.NewMap(&seed)
	return c.JSON(http.StatusOK, generation.GetTiles())
}
