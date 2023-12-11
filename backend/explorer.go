package theisland

// Can share tile with other explorers. Can move on the water or on the land.
// An explorer that have left the land cannot go back
// An explorer has 3 steps by lap. Each move to another land tile costs a step. Boarding or Landing on a boat costs a step.
// Swimming to another tile costs 3 steps
type Explorer struct {
	Id       int
	MaxSteps int
	Team     int
}

func NewExplorer(id int) *Explorer {
	return &Explorer{
		MaxSteps: 3,
		Team:     0,
		Id:       id,
	}
}
