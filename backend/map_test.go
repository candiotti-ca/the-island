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

func TestMapMoveBoat_unknownOriginTile(t *testing.T) {
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

func TestMapMoveBoat_noBoatOnTile(t *testing.T) {
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

func TestMapMoveBoat_unknownDestinationTile(t *testing.T) {
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

func TestMapMoveBoat_destinationTileIsNotWater(t *testing.T) {
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

func TestMapMoveBoatToTile_alreadyABoatOnDestinationTile(t *testing.T) {
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

func TestMapMoveBoat(t *testing.T) {
	t.Parallel()

	boat := NewBoat()
	orgCoord := Coord{Q: 1, R: 1}
	orgTile := Tile{Type: WATER, Boat: boat}
	destCoord := Coord{Q: 2, R: 2}
	destTile := Tile{Type: WATER}
	m := Map{
		tiles:           map[Coord]Tile{orgCoord: orgTile, destCoord: destTile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	result, err := m.MoveBoat(orgCoord, destCoord)
	require.NoError(t, err)
	assert.Equal(t, boat, result.Boat)

	orgTile = m.tiles[orgCoord]
	assert.Nil(t, orgTile.Boat)

	destTile = m.tiles[destCoord]
	assert.Equal(t, boat, destTile.Boat)
}

func TestMapMoveExplorer_unknownOriginTile(t *testing.T) {
	t.Parallel()

	m := Map{
		tiles:           map[Coord]Tile{},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.MoveExplorer(Coord{Q: 1, R: 1}, Coord{Q: 2, R: 2}, 1)
	require.ErrorContains(t, err, "origin coordinates outside the map")
}

func TestMapMoveExplorer_explorerNotOnTile(t *testing.T) {
	t.Parallel()

	orgCoord := Coord{Q: 1, R: 1}
	orgTile := NewTile(WATER)
	m := Map{
		tiles:           map[Coord]Tile{orgCoord: orgTile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.MoveExplorer(Coord{Q: 1, R: 1}, Coord{Q: 2, R: 2}, 1)
	require.ErrorContains(t, err, "explorer not present on this tile")
}

func TestMapMoveExplorer_unknownDestinationTile(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	orgCoord := Coord{Q: 1, R: 1}
	orgTile := Tile{Type: WATER, Explorers: []*Explorer{explorer}}
	m := Map{
		tiles:           map[Coord]Tile{orgCoord: orgTile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.MoveExplorer(orgCoord, Coord{Q: 2, R: 2}, explorer.Id)
	require.ErrorContains(t, err, "destination coordinates outside the map")
}

func TestMapMoveExplorer_moveFromWaterToLand(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	orgCoord := Coord{Q: 1, R: 1}
	orgTile := Tile{Type: WATER, Explorers: []*Explorer{explorer}}
	destCoord := Coord{Q: 2, R: 2}
	destTile := Tile{Type: ROCK}
	m := Map{
		tiles:           map[Coord]Tile{orgCoord: orgTile, destCoord: destTile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.MoveExplorer(orgCoord, destCoord, explorer.Id)
	require.ErrorContains(t, err, "a swimmer cannot go back to the land")
}

func TestMapMoveExplorer_sameTile(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	orgCoord := Coord{Q: 1, R: 1}
	orgTile := Tile{Type: ROCK, Explorers: []*Explorer{explorer}}
	m := Map{
		tiles:           map[Coord]Tile{orgCoord: orgTile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.MoveExplorer(orgCoord, orgCoord, explorer.Id)
	require.ErrorContains(t, err, "explorer already on the tile")
}

func TestMapMoveExplorer(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	orgCoord := Coord{Q: 1, R: 1}
	orgTile := Tile{Type: ROCK, Explorers: []*Explorer{explorer}}
	destCoord := Coord{Q: 2, R: 2}
	destTile := Tile{Type: WATER}
	m := Map{
		tiles:           map[Coord]Tile{orgCoord: orgTile, destCoord: destTile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	result, err := m.MoveExplorer(orgCoord, destCoord, explorer.Id)
	require.NoError(t, err)
	assert.NotNil(t, result.GetExplorer(explorer.Id))

	orgTile = m.tiles[orgCoord]
	assert.Nil(t, orgTile.GetExplorer(explorer.Id))

	destTile = m.tiles[destCoord]
	assert.NotNil(t, destTile.GetExplorer(explorer.Id))
}

func TestMapResolveInteractions_WhaleBoatExplorer(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	boat := NewBoat()
	err := boat.BoardExplorer(explorer)
	require.NoError(t, err)

	tile := NewTile(WATER)
	tile.Boat = boat
	tile.Whale = NewWhale()

	coord := Coord{Q: 1, R: 1}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	result, err := m.ResolveInteractions(coord)
	require.NoError(t, err)
	assert.NotNil(t, result.GetExplorer(explorer.Id))
	assert.NotNil(t, result.Whale)
	assert.Nil(t, result.Boat)
}

func TestMapResolveInteractions_SharkBoatExplorer(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	boat := NewBoat()
	err := boat.BoardExplorer(explorer)
	require.NoError(t, err)

	tile := NewTile(WATER)
	tile.Boat = boat
	tile.Shark = NewShark()

	coord := Coord{Q: 1, R: 1}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	result, err := m.ResolveInteractions(coord)
	require.NoError(t, err)
	assert.Nil(t, result.GetExplorer(explorer.Id))
	assert.NotNil(t, result.Shark)
	assert.NotNil(t, result.Boat)
	assert.Equal(t, 1, len(result.Boat.Explorers))
}

func TestMapResolveInteractions_SeaSerpentBoatExplorer(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	boat := NewBoat()
	err := boat.BoardExplorer(explorer)
	require.NoError(t, err)

	tile := NewTile(WATER)
	tile.Boat = boat
	tile.SeaSerpent = NewSeaSerpent()

	coord := Coord{Q: 1, R: 1}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	result, err := m.ResolveInteractions(coord)
	require.NoError(t, err)
	assert.Nil(t, result.GetExplorer(explorer.Id))
	assert.NotNil(t, result.SeaSerpent)
	assert.Nil(t, result.Boat)
}

func TestMapResolveInteractions_SeaSerpentExplorer(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	tile := NewTile(WATER)
	err := tile.AddExplorer(explorer)
	require.NoError(t, err)
	tile.SeaSerpent = NewSeaSerpent()

	coord := Coord{Q: 1, R: 1}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	result, err := m.ResolveInteractions(coord)
	require.NoError(t, err)
	assert.Nil(t, result.GetExplorer(explorer.Id))
	assert.NotNil(t, result.SeaSerpent)
}

func TestMapResolveInteractions_WhaleExplorer(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	tile := NewTile(WATER)
	err := tile.AddExplorer(explorer)
	require.NoError(t, err)
	tile.Whale = NewWhale()

	coord := Coord{Q: 1, R: 1}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	result, err := m.ResolveInteractions(coord)
	require.NoError(t, err)
	assert.NotNil(t, result.GetExplorer(explorer.Id))
	assert.NotNil(t, result.Whale)
}

func TestMapResolveInteractions_SharkExplorer(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	tile := NewTile(WATER)
	err := tile.AddExplorer(explorer)
	require.NoError(t, err)
	tile.Shark = NewShark()

	coord := Coord{Q: 1, R: 1}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	result, err := m.ResolveInteractions(coord)
	require.NoError(t, err)
	assert.Nil(t, result.GetExplorer(explorer.Id))
	assert.NotNil(t, result.Shark)
}

func TestMapResolveInteractions_emptyTile(t *testing.T) {
	t.Parallel()

	tile := NewTile(WATER)

	coord := Coord{Q: 1, R: 1}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	_, err := m.ResolveInteractions(coord)
	require.NoError(t, err)
}

func TestMapResolveInteractions_SeaSerpent(t *testing.T) {
	t.Parallel()

	tile := NewTile(WATER)
	tile.SeaSerpent = NewSeaSerpent()

	coord := Coord{Q: 1, R: 1}
	m := Map{
		tiles:           map[Coord]Tile{coord: tile},
		tileNumberRock:  8,
		tileNumberSand:  0,
		tileNumberGrass: 0,
	}

	result, err := m.ResolveInteractions(coord)
	require.NoError(t, err)
	assert.NotNil(t, result.SeaSerpent)
}
