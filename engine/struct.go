package engine

import (
	"github.con/enotinc/gotari/engine/cmd"
	"github.con/enotinc/gotari/engine/cursor"
	"github.con/enotinc/gotari/engine/options"
	"golang.org/x/term"
)

// main engine struct
type Engine struct {
	// list of games
	Games []*Game

	// calculates on Init func
	// used as starting point for render and to transfer real mouse click position to relative
	starting cursor.Position

	opt *options.Options

	// selected (currently played) game
	game *Game

	// game selection menu
	menu *Menu

	// map of line indexes to saved hashes
	hash map[int]uint32

	cursor *cursor.Cursor

	// additional stuff to work with terminal raw state
	fdIn int
	old  *term.State
}

type Menu interface {
	Render() []*string

	// Used to give Menu list of all games in enigne.Games
	LoadList(games []*Game)

	// if game was selected in menu, this game will be returned
	// if game wasn't selected - nil must me returned
	// CMD - usend only to check for QuitGotari. If it catch this CMD - programm will be closed
	Handle(event Event) cmd.CMD

	// if engine resieve a CMD.SelectGame - this function will be called to get selected game
	SelectedGame() *Game

	// calls before render
	CursorInfo() (cursor.Kind, cursor.Position)
}

type Game interface {
	// mainly used in menu, to get them name of the game
	Name() string

	// each string will be rendered with an engine 'diff render'.
	// You shoudn't add '\n\r' at the end of the line, engine will do this for you. And if you add - it will be replaced
	// And you shoudn't add '\n\r' anyware. This will most likely break render, coz it goes 'line by line'. See Engine.render() to see how it works
	Render() []*string

	// used to handle input keys. Mouse support isn't ready yet
	// func must return engine.CMD (CloseGame for example), or nil if nothing happened
	Handle(event Event) cmd.CMD

	// this funcion will be called before Render()
	// after cursor kind will be apllied
	CursorInfo() (cursor.Kind, cursor.Position)
}
