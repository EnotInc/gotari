package tictactoe

import (
	"github.con/enotinc/gotari/engine"
	"github.con/enotinc/gotari/enums/keys"
)

type sqare int

const (
	_ sqare = iota
	empty
	cross
	zero
)

type TTT struct {
	// board is an array of pointer to obj(booleans)
	// true  - cross
	// false - zero
	// nil   - empty
	board [9]sqare

	// additional message
	message string

	turn int // max 9
}

func (t *TTT) reset() {
	t.board = [9]sqare{empty, empty, empty, empty, empty, empty, empty, empty, empty}
	t.turn = 0
	t.message = ""
}

func Init() *TTT {
	t := &TTT{}
	t.reset()
	return t
}

func (t *TTT) CursorInfo() (engine.CursorKind, engine.CursorPosition) {
	return engine.Hidden, engine.CursorPosition{}
}

func (t *TTT) Render() []*string {
	return t.draw()
}

func (t *TTT) Handle(key rune) engine.CMD {
	switch key {
	case keys.Esc:
		return engine.CloseGame
	case 'r':
		t.reset()
		return nil
	}

	if t.isWin() != not {
		return nil
	}

	if ok := t.makeTurn(key); ok {
		t.turn += 1
		t.message = ""
	}
	return nil
}

func (t *TTT) Name() string {
	return "Tic-Tac-Toe"
}
