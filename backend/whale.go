package theisland

import (
	"github.com/google/uuid"
)

// Destroys a boat if it is occupied by explorers. Every explorer on the boat fall into the water. The whale does not hurt the explorers.
// Moves only in the water
type Whale struct {
	Id       uuid.UUID `json:"id"`
	MaxSteps int       `json:"maxSteps"`
}

func NewWhale() *Whale {
	return &Whale{
		Id:       uuid.New(),
		MaxSteps: 3,
	}
}
