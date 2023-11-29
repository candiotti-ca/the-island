package main

// Can share tile with other explorers. Can move on the water or on the land.
// An explorer that have left the land cannot go back
// An explorer has 3 steps by lap. Each move to another land tile costs a step. Boarding or Landing on a boat costs a step.
// Swimming to another tile costs 3 steps
type Explorer struct {
	Eliminated bool
	MaxSteps   int
	Team       int
}

func NewExplorer() *Explorer {
	return &Explorer{
		Eliminated: false,
		MaxSteps:   3,
		Team:       0,
	}
}
