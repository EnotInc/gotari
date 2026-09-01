package engine

import "golang.org/x/term"

// main engine struct
type Engine struct {
	// list of games
	Games []*Game

	// selected (currently played) game
	game *Game

	// game selection menu
	menu *Menu

	// map of line indexes to saved hashes
	hash map[int]uint32

	fullscreen bool

	// additional stuff to work with terminal raw state
	fdIn int
	old  *term.State
}

type Menu interface {
	Render() []*string
	LoadList(games []*Game)

	// if game was selected in menu, this game will be returned
	// if game wasn't selected - nil must me returned
	Handle(key rune) *Game
}

type Game interface {
	Name() string
	Render() []*string
	Handle(key rune)
}
