package theisland

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/falanger/hexgrid"
)

func hexCenter() hexgrid.Hex {
	return hexgrid.Hex{Q: 0, R: 0, S: 0}
}

type Layout hexgrid.Layout

type Map struct {
	// layout Layout
	tiles           map[Coord]Tile
	tileNumberRock  int
	tileNumberGrass int
	tileNumberSand  int
}

func (m Map) GetTiles() map[string]Tile {
	cp := make(map[string]Tile, 129)
	for coord, tile := range m.tiles {
		cp[coord.String()] = tile
	}
	return cp
}

func (m Map) GetTile(coord Coord) Tile {
	return m.tiles[coord]
}

func (m *Map) RemoveLandTile(coord Coord) (Tile, error) {
	tile := m.tiles[coord]

	switch tile.Type {
	case WATER:
		return Tile{}, errors.New("not a land tile")
	case SAND:
		m.tileNumberSand -= 1
	case GRASS:
		if m.tileNumberSand != 0 {
			return Tile{}, errors.New("remove sand tiles first")
		}
		m.tileNumberGrass -= 1
	case ROCK:
		if m.tileNumberSand != 0 && m.tileNumberRock != 0 {
			return Tile{}, errors.New("remove sand and grass tiles first")
		}
		m.tileNumberRock -= 1
	}

	tile.Type = WATER
	m.tiles[coord] = tile
	return tile, nil
}

type Coord struct {
	Q, R int
}

func (c Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.Q, c.R)
}

func NewMap(seed *int64) Map {
	var s int64
	if seed != nil {
		s = *seed
	} else {
		s = time.Now().Unix()
	}

	return Map{
		// layout: Layout{
		// 	Origin:      hexgrid.Point{X: 0, Y: 0},
		// 	Size:        hexgrid.Point{X: 15, Y: 15},
		// 	Orientation: hexgrid.OrientationPointy,
		// },
		tiles: initTiles(s),
	}
}

func initTiles(seed int64) map[Coord]Tile {
	rand := rand.New(rand.NewSource(seed))
	direction := hexgrid.Direction(rand.Intn(6))
	tiles := make(map[Coord]Tile, 129)

	fifthRing := sprialRing(direction, 5)
	tilesType := make([]TileType, 40)
	list := rand.Perm(40)
	for i, n := range list {
		if i < initTileNumberSand {
			tilesType[n] = SAND
		} else if i < initTileNumberSand+initTileNumberGrass {
			tilesType[n] = GRASS
		} else {
			tilesType[n] = ROCK
		}
	}
	for i, hex := range fifthRing[1:] {
		var tileType TileType
		if i < 40 {
			tileType = tilesType[i]
		} else {
			tileType = WATER
		}
		tiles[Coord{Q: hex.Q, R: hex.R}] = Tile{
			Type: tileType,
		}
	}

	sixthRing := singleRing(direction, 6)
	for _, hex := range sixthRing {
		if hex.Q == 0 && (hex.R == 6 || hex.R == -6) {
			continue
		}
		tiles[Coord{Q: hex.Q, R: hex.R}] = Tile{
			Type: WATER,
		}
	}

	tiles[Coord{Q: 4, R: 3}] = Tile{Type: WATER}
	tiles[Coord{Q: 5, R: 2}] = Tile{Type: WATER}
	tiles[Coord{Q: 4, R: -7}] = Tile{Type: WATER}
	tiles[Coord{Q: 5, R: -7}] = Tile{Type: WATER}
	tiles[Coord{Q: -4, R: -3}] = Tile{Type: WATER}
	tiles[Coord{Q: -5, R: -2}] = Tile{Type: WATER}
	tiles[Coord{Q: -4, R: -3}] = Tile{Type: WATER}
	tiles[Coord{Q: -5, R: -2}] = Tile{Type: WATER}

	return tiles
}

func singleRing(direction hexgrid.Direction, radius int) []hexgrid.Hex {
	nextDirection := func(dir hexgrid.Direction) hexgrid.Direction {
		switch dir {
		case hexgrid.DirectionN:
			return hexgrid.DirectionNW
		case hexgrid.DirectionNW:
			return hexgrid.DirectionSW
		case hexgrid.DirectionSW:
			return hexgrid.DirectionS
		case hexgrid.DirectionS:
			return hexgrid.DirectionSE
		case hexgrid.DirectionSE:
			return hexgrid.DirectionNE
		case hexgrid.DirectionNE:
			return hexgrid.DirectionN
		default:
			return hexgrid.DirectionN
		}
	}

	hex := hexCenter().Neighbor(direction).Scale(radius)
	currentDir := nextDirection(nextDirection(direction))

	result := make([]hexgrid.Hex, 0, 6*radius)
	for i := 0; i < 6; i++ {
		for j := 0; j < radius; j++ {
			result = append(result, hex)
			hex = hex.Neighbor(currentDir)
		}
		currentDir = nextDirection(currentDir)
	}
	return result
}

func sprialRing(direction hexgrid.Direction, radius int) []hexgrid.Hex {
	result := make([]hexgrid.Hex, 0, 1+3*radius*(radius+1))
	result = append(result, hexCenter())
	for i := 1; i <= radius; i++ {
		result = append(result, singleRing(direction, i)...)
	}
	return result
}

const (
	initTileNumberRock  = 8
	initTileNumberGrass = 16
	initTileNumberSand  = 16
)
