package theisland

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestLayoutCubeRing(t *testing.T) {
// 	t.Parallel()

// 	for k, v := range initTiles(1) {
// 		fmt.Printf("(%d,%d): %s\n", k.Q, k.R, v)
// 	}
// }

func TestMapCreation(t *testing.T) {
	t.Parallel()

	var seed int64 = time.Now().Unix()
	m := NewMap(&seed)

	rockCounter := 0
	sandCounter := 0
	grassCounter := 0

	for _, tile := range m.GetTiles() {
		if tile.Type == ROCK {
			rockCounter += 1
		}

		if tile.Type == SAND {
			sandCounter += 1
		}

		if tile.Type == GRASS {
			grassCounter += 1
		}
	}

	assert.Equal(t, rockCounter, initTileNumberRock)
	assert.Equal(t, sandCounter, initTileNumberSand)
	assert.Equal(t, grassCounter, initTileNumberGrass)
}

func TestMapRemoveLandTile_unknownTile(t *testing.T) {
	t.Parallel()

	coord := Coord{Q: 1, R: 1}
	m := Map{
		tiles:           map[Coord]Tile{},
		tileNumberRock:  0,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.RemoveLandTile(coord)
	require.ErrorContains(t, err, "coordinates outside the map")
}

func TestMapRemoveLandTile_waterTile(t *testing.T) {
	t.Parallel()

	coord := Coord{Q: 1, R: 1}
	tile := Tile{Type: WATER}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  0,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.RemoveLandTile(coord)
	require.ErrorContains(t, err, "not a land tile")
}

func TestMapRemoveLandTile_sandTile(t *testing.T) {
	t.Parallel()

	coord := Coord{Q: 1, R: 1}
	tile := Tile{Type: SAND}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  0,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	tile, err := m.RemoveLandTile(coord)
	require.NoError(t, err)
	assert.Equal(t, WATER, tile.Type)
}

func TestMapRemoveLandTile_grassTileWhereasSandRemaining(t *testing.T) {
	t.Parallel()

	coord := Coord{Q: 1, R: 1}
	tile := Tile{Type: GRASS}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  0,
		tileNumberSand:  1,
		tileNumberGrass: 0,
	}

	_, err := m.RemoveLandTile(coord)
	require.ErrorContains(t, err, "remove sand tiles first")
}

func TestMapRemoveLandTile_grassTile(t *testing.T) {
	t.Parallel()

	coord := Coord{Q: 1, R: 1}
	tile := Tile{Type: GRASS}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  1,
		tileNumberSand:  0,
		tileNumberGrass: 3,
	}

	tile, err := m.RemoveLandTile(coord)
	require.NoError(t, err)
	assert.Equal(t, WATER, tile.Type)
}

func TestMapRemoveLandTile_rockTileWhereasSandRemaining(t *testing.T) {
	t.Parallel()

	coord := Coord{Q: 1, R: 1}
	tile := Tile{Type: ROCK}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  0,
		tileNumberSand:  1,
		tileNumberGrass: 0,
	}

	_, err := m.RemoveLandTile(coord)
	require.ErrorContains(t, err, "remove sand and grass tiles first")
}

func TestMapRemoveLandTile_rockTileWhereasGrassRemaining(t *testing.T) {
	t.Parallel()

	coord := Coord{Q: 1, R: 1}
	tile := Tile{Type: ROCK}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 4,
	}

	_, err := m.RemoveLandTile(coord)
	require.ErrorContains(t, err, "remove sand and grass tiles first")
}

func TestMapRemoveLandTile_rockTile(t *testing.T) {
	t.Parallel()

	coord := Coord{Q: 1, R: 1}
	tile := Tile{Type: ROCK}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	tile, err := m.RemoveLandTile(coord)
	require.NoError(t, err)
	assert.Equal(t, WATER, tile.Type)
}

func TestBoatMoveToTile_unknownOriginTile(t *testing.T) {
	t.Parallel()

	m := Map{
		tiles:           map[Coord]Tile{},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.MoveBoat(Coord{Q: 1, R: 1}, Coord{Q: 2, R: 2})
	require.ErrorContains(t, err, "origin coordinates outside the map")
}

func TestBoatMoveToTile_noBoatOnTile(t *testing.T) {
	t.Parallel()

	orgCoord := Coord{Q: 1, R: 1}
	orgTile := Tile{Type: WATER}
	m := Map{
		tiles:           map[Coord]Tile{orgCoord: orgTile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.MoveBoat(orgCoord, Coord{Q: 2, R: 2})
	require.ErrorContains(t, err, "no boat on this tile")
}

func TestBoatMoveToTile_unknownDestinationTile(t *testing.T) {
	t.Parallel()

	boat := NewBoat()
	orgCoord := Coord{Q: 1, R: 1}
	orgTile := Tile{Type: WATER, Boat: boat}
	m := Map{
		tiles:           map[Coord]Tile{orgCoord: orgTile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.MoveBoat(orgCoord, Coord{Q: 2, R: 2})
	require.ErrorContains(t, err, "destination coordinates outside the map")
}

func TestBoatMoveToTile_destinationTileIsNotWater(t *testing.T) {
	t.Parallel()

	boat := NewBoat()
	orgCoord := Coord{Q: 1, R: 1}
	orgTile := Tile{Type: WATER, Boat: boat}
	destCoord := Coord{Q: 2, R: 2}
	destTile := Tile{Type: ROCK}
	m := Map{
		tiles:           map[Coord]Tile{orgCoord: orgTile, destCoord: destTile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.MoveBoat(orgCoord, destCoord)
	require.ErrorContains(t, err, "a boat can move only in water")
}

func TestBoatMoveToTile_alreadyABoatOnDestinationTile(t *testing.T) {
	t.Parallel()

	boat := NewBoat()
	orgCoord := Coord{Q: 1, R: 1}
	orgTile := Tile{Type: WATER, Boat: boat}
	destCoord := Coord{Q: 2, R: 2}
	destTile := Tile{Type: WATER, Boat: boat}
	m := Map{
		tiles:           map[Coord]Tile{orgCoord: orgTile, destCoord: destTile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.MoveBoat(orgCoord, destCoord)
	require.ErrorContains(t, err, "destination tile already has a boat")
}
