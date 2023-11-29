package main

import "fmt"

// Destroys a boat if it is occupied by explorers. Every explorer on the boat fall into the water. The whale does not hurt the explorers.
// Moves only in the water
type Whale struct {
	Eliminated bool
	MaxSteps   int
}

func NewWhale() *Whale {
	return &Whale{
		Eliminated: false,
		MaxSteps:   3,
	}
}

func (whale Whale) String() string {
	return fmt.Sprintf("Whale{Eliminated: %t, MaxSteps: %d}", whale.Eliminated, whale.MaxSteps)
}
