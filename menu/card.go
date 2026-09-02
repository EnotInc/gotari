package menu

import (
	"github.con/enotinc/gotari/engine"
)

type card struct {
	name string
	game *engine.Game
}

func NewCard(name string, game *engine.Game) *card {
	return &card{
		name: name,
		game: game,
	}
}
