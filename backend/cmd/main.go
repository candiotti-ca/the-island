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

	e.PATCH("/boat/:qf/:rf/move/:qt/:rt", moveBoat)
	e.PATCH("/shark/:qf/:rf/move/:qt/:rt", moveSth)
	e.PATCH("/sea_serpent/:qf/:rf/move/:qt/:rt", moveSth)
	e.PATCH("/whale/:qf/:rf/move/:qt/:rt", moveSth)
	e.PATCH("/explorer/:qf/:rf/move/:qt/:rt", moveSth)

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

func moveBoat(c echo.Context) error {
	Q, err := strconv.Atoi(c.Param("qf"))
	if err != nil {
		return errors.New("cannot get Q param")
	}

	R, err := strconv.Atoi(c.Param("rf"))
	if err != nil {
		return errors.New("cannot get R param")
	}

	from := theisland.Coord{Q, R}

	Q, err = strconv.Atoi(c.Param("qt"))
	if err != nil {
		return errors.New("cannot get Q param")
	}

	R, err = strconv.Atoi(c.Param("rt"))
	if err != nil {
		return errors.New("cannot get R param")
	}

	to := theisland.Coord{Q, R}

	destinationTile, err := GameMap.MoveBoat(from, to)
	if err != nil {
		return errors.New("cannot move boat")
	}

	return c.JSON(http.StatusOK, destinationTile)
}

func moveSth(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
