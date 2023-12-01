package theisland

type Tile struct {
	Type TileType `json:"type"`
}

type TileType int

const (
	WATER TileType = iota + 1
	SAND
	GRASS
	ROCK
)

func (t TileType) String() string {
	switch t {
	case SAND:
		return "SAND"
	case GRASS:
		return "GRASS"
	case ROCK:
		return "ROCK"
	}
	return "WATER"
}
