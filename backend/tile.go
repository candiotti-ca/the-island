package theisland

import (
	"errors"
	"slices"

	"github.com/google/uuid"
)

type Tile struct {
	Id         uuid.UUID   `json:"id"`
	Type       TileType    `json:"type"`
	Explorers  []*Explorer `json:"explorers"`
	Boat       *Boat       `json:"boat"`
	Shark      *Shark      `json:"shark"`
	Whale      *Whale      `json:"whale"`
	SeaSerpent *SeaSerpent `json:"seaSerpent"`
}

type Option func(o *Tile)

func NewTile(opts ...Option) Tile {
	tile := &Tile{
		Id:        uuid.New(),
		Type:      WATER,
		Explorers: make([]*Explorer, 0),
	}

	for _, opt := range opts {
		opt(tile)
	}

	return *tile
}

func WithTileType(tileType TileType) Option {
	return func(o *Tile) {
		o.Type = tileType
	}
}

func WithBoat() Option {
	return func(o *Tile) {
		o.Boat = NewBoat()
	}
}

func WithWhale() Option {
	return func(o *Tile) {
		o.Whale = NewWhale()
	}
}

func WithShark() Option {
	return func(o *Tile) {
		o.Shark = NewShark()
	}
}

func WithSeaSerpent() Option {
	return func(o *Tile) {
		o.SeaSerpent = NewSeaSerpent()
	}
}

func (t *Tile) RemoveExplorer(id int) (*Explorer, error) {
	index := -1
	var expToDelete *Explorer
	for i, explorer := range t.Explorers {
		if explorer.Id == id {
			index = i
			expToDelete = explorer
		}
	}

	if index == -1 {
		return nil, errors.New("explorer not present on the tile")
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

	return expToDelete, nil
}

func (t *Tile) AddExplorer(explorer *Explorer) error {
	if slices.Contains(t.Explorers, explorer) {
		return errors.New("explorer already on the tile")
	}

	t.Explorers = append(t.Explorers, explorer)
	return nil
}

func (t *Tile) GetExplorer(id int) *Explorer {
	for _, explorer := range t.Explorers {
		if explorer.Id == id {
			return explorer
		}
	}
	return nil
}

func (t *Tile) DestroyBoat() error {
	if t.Boat == nil {
		return nil
	}

	boat := t.Boat
	for _, explorer := range boat.Explorers {
		err := t.AddExplorer(explorer)
		if err != nil {
			return err
		}
	}
	t.Boat = nil

	return nil
}

func (t *Tile) KillSwimmers() {
	if t.Type == WATER {
		t.Explorers = make([]*Explorer, 0)
	}
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
