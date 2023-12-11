package theisland

import "fmt"

// Eats each explorer that are in the water.
// moves only in the water
type Shark struct {
	MaxSteps int
}

func NewShark() *Shark {
	return &Shark{
		MaxSteps: 2,
	}
}

func (shark Shark) String() string {
	return fmt.Sprintf("Shark{MaxSteps: %d}", shark.MaxSteps)
}
