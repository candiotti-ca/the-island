package main

type Tile struct {
	Type TileType
	X    int
	Y    int
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
