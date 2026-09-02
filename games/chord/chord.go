package chord

import (
	"time"

	"github.con/enotinc/gotari/engine"
	"github.con/enotinc/gotari/enums/keys"
)

type Chord struct {
	begin   time.Time
	end     time.Time
	typed   int // amount of typed symbols (backspace included)
	errors  int // amount of errors
	started bool
	ended   bool

	input  string // user input
	text   string // initial text
	cursor int    // cursor position
}

func (c *Chord) reset() {
	c.text = getWords()
	c.input = ""

	c.cursor = 0
	c.typed = 0
	c.errors = 0

	c.started = false
	c.ended = false
}

func Init() *Chord {
	c := &Chord{}
	c.reset()
	return c
}

func (c *Chord) CursorInfo() (engine.CursorKind, engine.CursorPosition) {
	// TODO: change cursor kind and position
	return engine.Hidden, engine.CursorPosition{}
}

func (c *Chord) Handle(key rune) engine.CMD {
	switch key {
	case keys.Esc:
		return engine.CloseGame
	default:
		c.handle(key)
	}
	return nil
}

func (c *Chord) Render() []*string {
	return c.render()
}

func (c *Chord) Name() string {
	return "Chord"
}
