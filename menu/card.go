package menu

import (
	"github.con/enotinc/gotari/engine"
)

const (
	horisontal  = "─"
	uppreCorner = "╮"
	vertical    = "│"
	lowerCorner = "╯"

	reset = "\033[0m"
	gray  = "\033[90m"
	blue  = "\033[34m"
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
