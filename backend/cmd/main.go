package main

import (
	"errors"
	"net/http"
	"strconv"
	"theisland"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var GameMap theisland.Map

func main() {
	var seed int64 = time.Now().Unix()
	GameMap = theisland.NewMap(&seed)

	e := echo.New()
	e.Use(middleware.CORS())
	// e.Use(middleware.Logger())

	e.GET("/map", generateMap)
	e.PATCH("/tiles/:q/:r/flip", flipTile)

	e.Logger.Fatal(e.Start(":8080"))
}

func generateMap(c echo.Context) error {
	return c.JSON(http.StatusOK, GameMap.GetTiles())
}

func flipTile(c echo.Context) error {
	Q, err := strconv.Atoi(c.Param("q"))
	if err != nil {
		return errors.New("cannot get Q param")
	}

	R, err := strconv.Atoi(c.Param("r"))
	if err != nil {
		return errors.New("cannot get R param")
	}

	flipedTile, err := GameMap.RemoveLandTile(theisland.Coord{Q: Q, R: R})
	if err != nil {
		return errors.New("cannot remove land tile")
	}

	return c.JSON(http.StatusOK, flipedTile)
}
