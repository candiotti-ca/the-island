package theisland

// Eats each explorer that are in the water
// moves only in the water
type Shark struct {
	Eliminated bool
	MaxSteps   int
}

func NewShark() *Shark {
	return &Shark{
		Eliminated: false,
		MaxSteps:   2,
	}
}
