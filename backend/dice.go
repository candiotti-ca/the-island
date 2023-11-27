package theisland

import "math/rand"

type Dice struct {
	faces int
}

func New() Dice {
	return Dice{faces: 6}
}

func (d Dice) Throw() int {
	return rand.Intn(d.faces-1) + 1
}
