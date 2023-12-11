package theisland

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTileRemoveExplorer_noExplorer(t *testing.T) {
	t.Parallel()

	explorer1 := NewExplorer(1)
	explorer2 := NewExplorer(2)
	tile := NewTile(WATER)

	err := tile.AddExplorer(explorer1)
	require.NoError(t, err)
	deleted, err := tile.RemoveExplorer(explorer2.Id)

	require.ErrorContains(t, err, "explorer not present on the tile")
	assert.Nil(t, deleted)
	assert.Equal(t, 1, len(tile.Explorers))
}

func TestTileRemoveExplorer_oneExplorer(t *testing.T) {
	t.Parallel()

	explorer := NewExplorer(1)
	tile := NewTile(WATER)

	err := tile.AddExplorer(explorer)
	require.NoError(t, err)
	deleted, err := tile.RemoveExplorer(explorer.Id)

	require.NoError(t, err)
	assert.Equal(t, explorer, deleted)
	assert.Equal(t, 0, len(tile.Explorers))
}

func TestTileRemoveExplorer_twoExplorers(t *testing.T) {
	t.Parallel()

	explorer1 := NewExplorer(1)
	explorer2 := NewExplorer(2)
	tile := NewTile(WATER)

	err := tile.AddExplorer(explorer1)
	require.NoError(t, err)
	err = tile.AddExplorer(explorer2)
	require.NoError(t, err)
	deleted, err := tile.RemoveExplorer(explorer2.Id)

	require.NoError(t, err)
	assert.Equal(t, explorer2, deleted)
	assert.Equal(t, 1, len(tile.Explorers))
}

func TestTileGetExplorer(t *testing.T) {
	t.Parallel()

	explorer1 := NewExplorer(1)
	explorer2 := NewExplorer(2)
	tile := NewTile(WATER)

	err := tile.AddExplorer(explorer1)
	require.NoError(t, err)
	assert.Equal(t, explorer1, tile.GetExplorer(explorer1.Id))
	assert.Nil(t, tile.GetExplorer(explorer2.Id))
}

func TestTileAddExplorer(t *testing.T) {
	t.Parallel()

	explorer1 := NewExplorer(1)

	tile := NewTile(WATER)
	assert.Equal(t, 0, len(tile.Explorers))

	err := tile.AddExplorer(explorer1)
	require.NoError(t, err)
	assert.Equal(t, 1, len(tile.Explorers))
	assert.NotNil(t, tile.GetExplorer(explorer1.Id))

	err = tile.AddExplorer(explorer1)
	require.ErrorContains(t, err, "explorer already on the tile")
	assert.Equal(t, 1, len(tile.Explorers))
}

func TestTileDestroyBoat_noBoat(t *testing.T) {
	t.Parallel()

	tile := NewTile(WATER)

	err := tile.DestroyBoat()
	require.NoError(t, err)
}

func TestTileDestroyBoat_withoutExplorers(t *testing.T) {
	t.Parallel()

	tile := NewTile(WATER)
	boat := NewBoat()
	tile.Boat = boat

	err := tile.DestroyBoat()
	require.NoError(t, err)
}

func TestTileDestroyBoat_withExplorers(t *testing.T) {
	t.Parallel()

	tile := NewTile(WATER)
	boat := NewBoat()
	boat.Explorers = append(boat.Explorers, NewExplorer(1))
	tile.Boat = boat

	err := tile.DestroyBoat()
	require.NoError(t, err)
	assert.Equal(t, 1, len(tile.Explorers))
}

func TestTileKillSwimmers_landTile(t *testing.T) {
	t.Parallel()

	tile := NewTile(ROCK)
	err := tile.AddExplorer(NewExplorer(1))
	assert.NoError(t, err)

	tile.KillSwimmers()
	assert.Equal(t, 1, len(tile.Explorers))
}

func TestTileKillSwimmers(t *testing.T) {
	t.Parallel()

	tile := NewTile(WATER)
	err := tile.AddExplorer(NewExplorer(1))
	assert.NoError(t, err)

	tile.KillSwimmers()
	assert.Equal(t, 0, len(tile.Explorers))
}
