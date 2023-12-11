package theisland

type Tile struct {
	Type       TileType    `json:"type"`
	Explorers  []*Explorer `json:"explorers"`
	Boat       *Boat       `json:"boat"`
	Shark      *Shark      `json:"shark"`
	Whale      *Whale      `json:"whale"`
	SeaSerpent *SeaSerpent `json:"seaSerpent"`
}

func NewTile(tileType TileType) Tile {
	return Tile{Type: tileType, Explorers: make([]*Explorer, 0)}
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
