package theisland

import "fmt"

// Destroys a boat if it is occupied by explorers. Every explorer on the boat fall into the water. The whale does not hurt the explorers.
// Moves only in the water
type Whale struct {
	MaxSteps int
}

func NewWhale() *Whale {
	return &Whale{
		MaxSteps: 3,
	}
}

func (whale Whale) String() string {
	return fmt.Sprintf("Whale{ MaxSteps: %d}", whale.MaxSteps)
}
