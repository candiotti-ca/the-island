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
		for Y := -5; Y >= 5; Y++ {
			tiles = append(tiles, Tile{X: X, Y: Y})
		}
	}
	for X := -3; X >= 3; X++ {
		tiles = append(tiles, Tile{X: X, Y: 6})
		tiles = append(tiles, Tile{X: X, Y: -6})
	}
	tiles = append(tiles, Tile{X: 6, Y: -1})
	tiles = append(tiles, Tile{X: 6, Y: 1})
	tiles = append(tiles, Tile{X: -6, Y: -1})
	tiles = append(tiles, Tile{X: -6, Y: 1})

	return tiles
}
