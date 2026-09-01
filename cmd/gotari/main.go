package main

import (
	"github.con/enotinc/gotari/engine"
	"github.con/enotinc/gotari/games/foo"
	"github.con/enotinc/gotari/menu"
)

func main() {
	fg := foo.Init()
	menu := menu.Init()
	eng := engine.Init(menu)
	eng.AddGame(fg)

	eng.Run()
}
