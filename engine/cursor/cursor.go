package cursor

import (
	"fmt"
)

type Position struct {
	X int
	Y int
}

type Cursor struct {
	kind Kind

	// Position is applyed to the begining of the render
	// on fullscreen - starting point is top left of terminal window
	// otherwise - starting point is at the begining of the line, right below the prompt
	pos Position
}

type Kind int

const (
	_ Kind = iota
	Block
	Line
	Underline
	Hidden
)

const (
	_hideCursor = "\033[?25l"
	_showCursor = "\033[?25h"

	_resetCursor     = "\033[0 q"
	_blockCursor     = "\033[2 q"
	_lineCursor      = "\033[6 q"
	_underlineCursor = "\033[4 q"
)

func (c *Cursor) Reset() string {
	return fmt.Sprintf("%s%s", _showCursor, _resetCursor)
}

func NewCursor() *Cursor {
	return &Cursor{}
}

func (c *Cursor) ChangePos(p Position) {
	if p != (Position{}) {
		c.pos = p
	}
}

func (c *Cursor) ChangeCursor(new Kind) {
	if c.kind == new {
		return // to nothing if new cursor type is the same as current one
	}

	if c.kind == Hidden { // if currend cursor is hidden: show it
		fmt.Print(_showCursor)
	}

	switch new { // apply new cursor kind
	case Hidden:
		fmt.Print(_hideCursor)
	case Block:
		fmt.Print(_blockCursor)
	case Line:
		fmt.Print(_lineCursor)
	case Underline:
		fmt.Print(_underlineCursor)
	}

	c.kind = new
}

func (c *Cursor) GetRealPos(starting Position) (x int, y int) {
	return starting.X + c.pos.X, starting.Y + c.pos.Y
}
