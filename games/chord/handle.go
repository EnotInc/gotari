package chord

import (
	"fmt"
	"time"

	"github.con/enotinc/gotari/enums/keys"
)

func isKey(key rune) bool {
	return ('a' <= key && key <= 'z') || ('A' <= key && key <= 'Z') || key == keys.Space || key == keys.Backspace
}

func (c *Chord) handle(key rune) {
	if !isKey(key) {
		return
	}

	if !c.started {
		c.started = true
		c.begin = time.Now()
	}

	switch key {
	case keys.Backspace:
		if c.cursor > 0 {
			c.cursor -= 1
			c.input = c.input[:c.cursor]
		}
	default:
		if c.ended {
			if key == 'r' {
				c.reset()
			}
			return
		}

		c.typed += 1
		if key != rune(c.text[c.cursor]) {
			c.errors += 1
		}

		c.input = fmt.Sprintf("%s%c", c.input, key)
		c.cursor += 1
	}

	if c.cursor == len(c.text) {
		c.ended = true
		c.end = time.Now()
	}
}
