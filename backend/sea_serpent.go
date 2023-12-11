package theisland

import "fmt"

// Destroys a boat if it was occupied by explorers. Eats each explorer in the water or on the boat.
// Moves only in the water
type SeaSerpent struct {
	MaxSteps int
}

func NewSeaSerpent() *SeaSerpent {
	return &SeaSerpent{
		MaxSteps: 1,
	}
}

func (seaSerpent SeaSerpent) String() string {
	return fmt.Sprintf("SeaSerpent{MaxSteps: %d}", seaSerpent.MaxSteps)
}
