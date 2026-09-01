package main

import (
	"github.con/enotinc/gotari/engine"
	"github.con/enotinc/gotari/games/chord"
	tictactoe "github.con/enotinc/gotari/games/tic-tac-toe"
	"github.con/enotinc/gotari/menu"
)

func main() {
	menu := menu.Init()
	eng := engine.Init(menu)

	ttt := tictactoe.Init()
	chord := chord.Init()

	eng.AddGame(ttt)
	eng.AddGame(chord)

	eng.Run()
}
