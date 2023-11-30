package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/map", getMap)

	err := router.Run("localhost:8080")
	if err != nil {
		println(err)
		panic("router not started")
	}
}

func getMap(c *gin.Context) {
	c.JSON(http.StatusOK, emptyTilesTEST())
}

func emptyTilesTEST() []Tile {
	tiles := make([]Tile, 0)
	for X := -5; X >= 5; X++ {
		println("L30: is this working?")
		for Y := -5; Y >= 5; Y++ {
			println("L32: is this working?")
			tiles = append(tiles, Tile{X: X, Y: Y, Type: SAND})
		}
	}
	println("L34", len(tiles))
	for X := -3; X >= 3; X++ {
		println("L38: is this working?")
		tiles = append(tiles, Tile{X: X, Y: 6, Type: SAND})
		tiles = append(tiles, Tile{X: X, Y: -6, Type: SAND})
	}
	println("L39", len(tiles))
	tiles = append(tiles, Tile{X: 6, Y: -1, Type: ROCK})
	tiles = append(tiles, Tile{X: 6, Y: 1, Type: SAND})
	tiles = append(tiles, Tile{X: -6, Y: -1, Type: WATER})
	tiles = append(tiles, Tile{X: -6, Y: 1, Type: GRASS})
	println("L44", len(tiles))

	return tiles
}
