package theisland

// Destroy a boat if it was occupied by explorers. Eats each explorer in the water or on the boat.
// moves only in the water
type SeaSerpent struct {
	Eliminated bool
	MaxSteps   int
}

func NewSeaSerpent() *SeaSerpent {
	return &SeaSerpent{
		Eliminated: false,
		MaxSteps:   1,
	}
}
