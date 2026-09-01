package engine

import (
	"fmt"
	"strings"

	"github.con/enotinc/gotari/enums/ascii"
)

func (e *Engine) render() {
	if e.game == nil {
		e.renderMenu()
	} else {
		e.renderGame()
	}
}

func (e *Engine) renderMenu() {
	var diff strings.Builder

	render := (*e.menu).Render()

	if !e.fullscreen {
		diff.WriteString(ascii.ResetCursorPos)
	}

	for index, line := range render {
		curHash := getHash(*line)
		oldHash, ok := e.hash[index]

		if !ok || (ok && curHash != oldHash) {
			var pos string

			if e.fullscreen {
				pos = fmt.Sprintf("\033[%d;%d;H", index+termOffset, termOffset)
			} else {
				pos = "\n"
			}

			clear := "\033[0K"

			diff.WriteString(pos)
			diff.WriteString(clear)
			diff.WriteString(*line)
		}
	}
	fmt.Print(diff.String())
}

func (e *Engine) renderGame() {
	var diff strings.Builder

	render := (*e.game).Render()

	if !e.fullscreen {
		diff.WriteString(ascii.ResetCursorPos)
	}

	for index, line := range render {
		curHash := getHash(*line)
		oldHash, ok := e.hash[index]

		if !ok || (ok && curHash != oldHash) {
			var pos string

			if e.fullscreen {
				pos = fmt.Sprintf("\033[%d;%d;H", index+termOffset, termOffset)
			} else {
				pos = "\n"
			}

			clear := "\033[0K"

			diff.WriteString(pos)
			diff.WriteString(clear)
			diff.WriteString(*line)
		}
	}
	fmt.Print(diff.String())
}
