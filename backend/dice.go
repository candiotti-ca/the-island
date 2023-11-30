package theisland

import (
	"math/rand"
)

type Dice struct {
	faces int
}

func NewDice() Dice {
	return Dice{faces: 6}
}

func (d Dice) Throw() int {
	return rand.Intn(d.faces)
}
