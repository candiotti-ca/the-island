package theisland

import (
	"fmt"
	"testing"
)

func TestLayoutCubeRing(t *testing.T) {
	t.Parallel()

	for k, v := range initTiles(1) {
		fmt.Printf("(%d,%d): %s\n", k.Q, k.R, v)
	}
}
