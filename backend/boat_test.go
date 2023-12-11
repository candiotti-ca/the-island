package theisland

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoatBoardExplorer_explorerAlreadyOnBoat(t *testing.T) {
	t.Parallel()

	boat := NewBoat()
	explorer := NewExplorer()

	assert.Len(t, boat.Explorers, 0)

	err := boat.BoardExplorer(explorer)
	require.NoError(t, err)
	assert.Len(t, boat.Explorers, 1)

	err = boat.BoardExplorer(explorer)
	require.ErrorContains(t, err, "explorer is already on the boat")
	assert.Len(t, boat.Explorers, 1)

	assert.Equal(t, boat.Explorers[0], explorer)
}

func TestBoatBoardExplorer(t *testing.T) {
	t.Parallel()

	boat := NewBoat()
	explorer := NewExplorer()

	assert.Len(t, boat.Explorers, 0)

	err := boat.BoardExplorer(explorer)
	require.NoError(t, err)
	assert.Len(t, boat.Explorers, 1)
	assert.Equal(t, boat.Explorers[0], explorer)
}

func TestBoatBoardExplorer_boatIsFull(t *testing.T) {
	t.Parallel()

	boat := NewBoat()

	explorer1 := NewExplorer()
	err := boat.BoardExplorer(explorer1)
	require.NoError(t, err)
	assert.Len(t, boat.Explorers, 1)

	explorer2 := NewExplorer()
	err = boat.BoardExplorer(explorer2)
	require.NoError(t, err)
	assert.Len(t, boat.Explorers, 2)

	explorer3 := NewExplorer()
	err = boat.BoardExplorer(explorer3)
	require.NoError(t, err)
	assert.Len(t, boat.Explorers, 3)

	explorer4 := NewExplorer()
	err = boat.BoardExplorer(explorer4)
	require.ErrorContains(t, err, "boat is full")

	assert.Equal(t, boat.Explorers[0], explorer1)
	assert.Equal(t, boat.Explorers[1], explorer2)
	assert.Equal(t, boat.Explorers[2], explorer3)
}

func TestBoatLandExplorer(t *testing.T) {
	t.Parallel()

	boat := NewBoat()
	explorer := NewExplorer()

	err := boat.BoardExplorer(explorer)
	require.NoError(t, err)

	err = boat.LandExplorer(explorer)
	require.NoError(t, err)

	assert.Len(t, boat.Explorers, 0)
}

func TestBoatLandExplorer_boatIsEmpty(t *testing.T) {
	boat := NewBoat()
	explorer := NewExplorer()

	err := boat.LandExplorer(explorer)
	require.Error(t, err, "explorer is not on the boat")
}

func TestBoatLandExplorer_explorerNotOnBoat(t *testing.T) {
	t.Parallel()

	boat := NewBoat()
	explorer1 := NewExplorer()
	explorer2 := NewExplorer()

	err := boat.BoardExplorer(explorer1)
	require.NoError(t, err)

	err = boat.LandExplorer(explorer2)
	require.ErrorContains(t, err, "explorer is not on the boat")
	assert.Len(t, boat.Explorers, 1)
	assert.Equal(t, boat.Explorers[0], explorer1)
}

func TestBoatExplorerIndex_explorerOnBoat(t *testing.T) {
	t.Parallel()

	boat := NewBoat()
	explorer1 := NewExplorer()
	explorer2 := NewExplorer()

	boat.Explorers = []*Explorer{
		explorer1,
		explorer2,
	}

	assert.Equal(t, boat.ExplorerIndex(explorer1), 0)
	assert.Equal(t, boat.ExplorerIndex(explorer2), 1)
}

func TestBoatExplorerIndex_explorerNotOnBoat(t *testing.T) {
	t.Parallel()

	boat := NewBoat()
	explorer1 := NewExplorer()
	explorer2 := NewExplorer()

	assert.Equal(t, boat.ExplorerIndex(explorer1), -1)

	boat.Explorers = []*Explorer{
		explorer1,
	}

	assert.Equal(t, boat.ExplorerIndex(explorer2), -1)
}
