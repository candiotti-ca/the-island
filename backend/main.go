package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/map", generateMap)

	e.Logger.Fatal(e.Start(":8080"))
}

func generateMap(c echo.Context) error {
	return c.JSON(http.StatusOK, emptyTilesTEST())
}

func emptyTilesTEST() []Tile {
	tiles := make([]Tile, 0)
	for X := -5; X <= 5; X++ {
		for Y := -5; Y <= 5; Y++ {
			tiles = append(tiles, Tile{X: X, Y: Y, Type: ROCK})
		}
	}

	for X := -3; X <= 3; X++ {
		tiles = append(tiles, Tile{X: X, Y: 6, Type: GRASS})
		tiles = append(tiles, Tile{X: X, Y: -6, Type: SAND})
	}

	tiles = append(tiles, Tile{X: 6, Y: -1, Type: ROCK})
	tiles = append(tiles, Tile{X: 6, Y: 1, Type: SAND})
	tiles = append(tiles, Tile{X: -6, Y: -1, Type: WATER})
	tiles = append(tiles, Tile{X: -6, Y: 1, Type: GRASS})

	return tiles
}
