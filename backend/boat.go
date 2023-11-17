package theisland

import "errors"

// 1 boat per tile at a time but can share a tile with monsters or explorers.
// moves only in the water
type Boat struct {
	Eliminated bool
	MaxSteps   int
	Explorers  []*Explorer
}

func NewBoat() *Boat {
	return &Boat{
		Eliminated: false,
		MaxSteps:   3,
		Explorers:  make([]*Explorer, 0, 3),
	}
}

func (boat *Boat) BoardExplorer(explorer *Explorer) error {
	passengers := len(boat.Explorers)
	if passengers == 3 {
		return errors.New("Boat is full")
	}

	boat.Explorers[passengers-1] = explorer
	return nil
}

func (boat *Boat) LandExplorer(explorer *Explorer) error {
	explorerIndex := -1
	for index, passenger := range boat.Explorers {
		if passenger == explorer {
			explorerIndex = index
		}
	}

	if explorerIndex == -1 {
		return errors.New("Explorer is not on the boat")
	}

	boat.Explorers = append(boat.Explorers[:explorerIndex], boat.Explorers[explorerIndex+1:]...)
	return nil
}
