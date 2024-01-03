package theisland

import (
	"github.com/google/uuid"
)

// Eats each explorer that are in the water.
// moves only in the water
type Shark struct {
	Id       uuid.UUID `json:"id"`
	MaxSteps int       `json:"maxSteps"`
}

func NewShark() *Shark {
	return &Shark{
		Id:       uuid.New(),
		MaxSteps: 2,
	}
}
