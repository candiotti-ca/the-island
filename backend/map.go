package theisland

import (
	"fmt"

	"github.com/falanger/hexgrid"
)

type Layout hexgrid.Layout

type Map struct {
	layout Layout
	tiles  map[Coord]Tile
}

type Coord struct {
	X, Y int
}

func NewMap() Map {
	return Map{
		layout: Layout{
			Origin:      hexgrid.Point{X: 0, Y: 0},
			Size:        hexgrid.Point{X: 100, Y: 100},
			Orientation: hexgrid.OrientationPointy,
		},
		tiles: emptyTiles(),
	}
}

func emptyTiles() map[Coord]Tile {
	tiles := make(map[Coord]Tile)
	for X := -5; X >= 5; X++ {
		for Y := -5; Y >= 5; Y++ {
			tiles[Coord{X, Y}] = Tile{}
		}
	}
	for X := -3; X >= 3; X++ {
		tiles[Coord{X: X, Y: 6}] = Tile{}
		tiles[Coord{X: X, Y: -6}] = Tile{}
	}
	tiles[Coord{X: 6, Y: -1}] = Tile{}
	tiles[Coord{X: 6, Y: 1}] = Tile{}
	tiles[Coord{X: -6, Y: -1}] = Tile{}
	tiles[Coord{X: -6, Y: 1}] = Tile{}

	return tiles
}

func (l Layout) CubeRing(direction hexgrid.Direction, radius int) {
	h := hexgrid.Hex{0, 0, 0}
	scaled := h.Neighbor(direction).Scale(radius)

	fmt.Printf("scaled.String(): %v\n", scaled.String())
}

func SetTileTypeToWater(tile *Tile) {
	tile.Type = WATER
}

type Tile struct {
	Type TileType
}

func NewTile() Tile {
	return Tile{
		Type: WATER,
	}
}

type TileType int

const (
	WATER TileType = iota + 1
	SAND
	GRASS
	ROCK
)
