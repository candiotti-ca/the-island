package theisland

// Can share tile with other explorers. Can move on the water or on the land.
// An explorer that left land cannot go back
// An explorer has 3 steps by lap. Each move to another land tile costs a step. Boarding or Landing on a boat costs a step.
// Swimming to another tile costs 3 steps
type Explorer struct {
	Eliminated bool
	MaxSteps   int
}
