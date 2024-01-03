package theisland

import (
	"errors"

	"github.com/google/uuid"
)

// 1 boat per tile at a time but can share a tile with monsters or explorers.
// Moves only in the water
type Boat struct {
	Id        uuid.UUID   `json:"id"`
	MaxSteps  int         `json:"maxSteps"`
	Explorers []*Explorer `json:"explorers"`
}

func NewBoat() *Boat {
	return &Boat{
		Id:        uuid.New(),
		MaxSteps:  3,
		Explorers: make([]*Explorer, 0, 3),
	}
}

func (boat *Boat) BoardExplorer(explorer *Explorer) error {
	if len(boat.Explorers) == 3 {
		return errors.New("boat is full")
	}

	if boat.ExplorerIndex(explorer) != -1 {
		return errors.New("explorer is already on the boat")
	}

	boat.Explorers = append(boat.Explorers, explorer)
	return nil
}

func (boat *Boat) LandExplorer(explorer *Explorer) error {
	explorerIndex := boat.ExplorerIndex(explorer)

	if explorerIndex == -1 {
		return errors.New("explorer is not on the boat")
	}

	boat.Explorers = append(boat.Explorers[:explorerIndex], boat.Explorers[explorerIndex+1:]...)
	return nil
}

func (boat *Boat) ExplorerIndex(explorer *Explorer) int {
	explorerIndex := -1
	for index, passenger := range boat.Explorers {
		if passenger == explorer {
			explorerIndex = index
		}
	}

	return explorerIndex
}
