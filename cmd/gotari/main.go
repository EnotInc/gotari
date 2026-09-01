package main

import (
	"github.con/enotinc/gotari/engine"
	"github.con/enotinc/gotari/games/foo"
	tictactoe "github.con/enotinc/gotari/games/tic-tac-toe"
	"github.con/enotinc/gotari/menu"
)

func main() {
	fg := foo.Init()
	ttt := tictactoe.Init()
	menu := menu.Init()
	eng := engine.Init(menu)
	eng.AddGame(fg)
	eng.AddGame(ttt)

	eng.Run()
}
