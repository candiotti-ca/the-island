package theisland

import (
	"github.com/google/uuid"
)

// Destroys a boat if it was occupied by explorers. Eats each explorer in the water or on the boat.
// Moves only in the water
type SeaSerpent struct {
	Id       uuid.UUID `json:"id"`
	MaxSteps int       `json:"maxSteps"`
}

func NewSeaSerpent() *SeaSerpent {
	return &SeaSerpent{
		Id:       uuid.New(),
		MaxSteps: 1,
	}
}
