package theisland

import (
	"testing"

	"github.com/falanger/hexgrid"
)

func TestLayoutCubeRing(t *testing.T) {
	t.Parallel()

	m := NewMap()

	m.layout.CubeRing(hexgrid.DirectionN, 3)
}
