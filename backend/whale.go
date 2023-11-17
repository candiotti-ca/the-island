package theisland

// Destroy a boat if it is occupied by explorers. Every explorer on the boat fall into the water. The whale does not hurt the explorers.
// moves only in the water
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
