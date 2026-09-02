package engine

import "fmt"

type CursorKind int

// TODO: move to enigne/cursor
const (
	_ CursorKind = iota
	Block
	Line
	Underline
	Hidden
)

const (
	_hideCursor = "\033[?25l"
	_showCursor = "\033[?25h"

	_blockCursor     = "\033[2 q"
	_lineCursor      = "\033[6 q"
	_underlineCursor = "\033[4 q"
)

func newCursor() *cursor {
	return &cursor{}
}

func (c *cursor) changePos(p Position) {
	c.pos = p
}

func (c *cursor) changeCursor(new CursorKind) {
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

func (c cursor) getRealPos(starting Position) (x int, y int) {
	return starting.X + c.pos.X, starting.Y + c.pos.Y
}
