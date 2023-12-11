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

func (t *Tile) RemoveExplorer(id int) {
	index := -1
	for i, explorer := range t.Explorers {
		if explorer.Id == id {
			index = i
		}
	}

	if index > -1 {
		length := len(t.Explorers)
		if length == 1 {
			t.Explorers = make([]*Explorer, 0)
		} else {
			t.Explorers[index] = t.Explorers[length-1]
			t.Explorers = t.Explorers[:length-1]
		}
	}
}

func (t *Tile) AddExplorer(explorer *Explorer) {
	t.Explorers = append(t.Explorers, explorer)
}

func (t *Tile) GetExplorer(id int) *Explorer {
	for _, explorer := range t.Explorers {
		if explorer.Id == id {
			return explorer
		}
	}
	return nil
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
