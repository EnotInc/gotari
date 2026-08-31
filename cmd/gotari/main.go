package main

import (
	"github.con/enotinc/gotari/engine"
	"github.con/enotinc/gotari/games/foo"
)

func main() {
	fg := foo.Init()
	eng := engine.Init()
	eng.AddGame(fg)

	eng.Run()
}
