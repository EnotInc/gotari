package engine

import (
	"github.con/enotinc/gotari/engine/cmd"
	"golang.org/x/term"
)

// main engine struct
type Engine struct {
	// list of games
	Games []*Game

	starting Position

	// selected (currently played) game
	game *Game

	// game selection menu
	menu *Menu

	// map of line indexes to saved hashes
	hash map[int]uint32

	fullscreen bool
	cursor     *cursor

	// additional stuff to work with terminal raw state
	fdIn int
	old  *term.State
}

type Position struct {
	X,
	Y int
}

type cursor struct {
	// kind must me changed with cursor.changeKind() func
	// this way cursor will be hide and showed properly
	// otherwise new cursor kind will not be applyed
	kind CursorKind

	// Position is applyed to the begining of the render
	// on fullscreen - starting point is top left of terminal window
	// otherwise - starting point is at the begining of the line, right below the prompt
	pos Position
}

type Menu interface {
	Render() []*string

	// Used to give Menu list of all games in enigne.Games
	LoadList(games []*Game)

	// if game was selected in menu, this game will be returned
	// if game wasn't selected - nil must me returned
	// CMD - usend only to check for QuitGotari. If it catch this CMD - programm will be closed
	Handle(key rune) cmd.CMD

	// if engine resieve a CMD.SelectGame - this function will be called to get selected game
	SelectedGame() *Game
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
	Handle(key rune) cmd.CMD

	// this funcion will be called before Render()
	// after cursor kind will be apllied
	CursorInfo() (CursorKind, Position)
}
