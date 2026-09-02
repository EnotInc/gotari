package engine

import (
	"fmt"
	"strings"

	"github.con/enotinc/gotari/enums/ascii"
)

func (e *Engine) render() {
	if e.game == nil {
		kind, pos := (*e.menu).CursorInfo()
		e.cursor.ChangeCursor(kind)
		e.cursor.ChangePos(pos)

		e.renderMenu()
	} else {
		kind, pos := (*e.game).CursorInfo()
		e.cursor.ChangeCursor(kind)
		e.cursor.ChangePos(pos)

		e.renderGame()
	}
}

func trimNewLine(l string) string {
	var trim = l
	trim = strings.TrimSuffix(trim, "\n\r")
	trim = strings.TrimSuffix(trim, "\n")
	trim = strings.TrimSuffix(trim, "\r")
	trim = strings.TrimPrefix(trim, "\n\r")
	trim = strings.TrimPrefix(trim, "\n")
	trim = strings.TrimPrefix(trim, "\r")
	return trim
}

func (e *Engine) clear() {
	var clear strings.Builder
	if e.fullscreen {
		clear.WriteString(ascii.ClearView)
		clear.WriteString(ascii.MoveToStart)
	} else {
		clear.WriteString(ascii.ResetCursorPos)
		clear.WriteString(ascii.MoveDown)
		clear.WriteString(ascii.ClearAfter)
	}
	fmt.Print(clear.String())
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
			trim := trimNewLine(*line)

			if e.fullscreen {
				pos = fmt.Sprintf("\033[%d;%d;H", index+termOffset, termOffset)
			} else {
				pos = "\n\r"
			}

			clear := "\033[0K"

			diff.WriteString(pos)
			diff.WriteString(clear)
			diff.WriteString(trim)
		}
	}
	x, y := e.cursor.GetRealPos(e.starting)
	moveto := fmt.Sprintf("\033[%d;%dH", y, x)
	diff.WriteString(moveto)

	fmt.Print(diff.String())
}

func (e *Engine) renderGame() {
	var diff strings.Builder

	kind, _ := (*e.game).CursorInfo()
	e.cursor.ChangeCursor(kind)
	render := (*e.game).Render()

	if !e.fullscreen {
		diff.WriteString(ascii.ResetCursorPos)
	}

	for index, line := range render {
		curHash := getHash(*line)
		oldHash, ok := e.hash[index]

		if !ok || (ok && curHash != oldHash) {
			var pos string
			trim := trimNewLine(*line)

			if e.fullscreen {
				pos = fmt.Sprintf("\033[%d;%d;H", index+termOffset, termOffset)
			} else {
				pos = "\n\r"
			}

			clear := "\033[0K"

			diff.WriteString(pos)
			diff.WriteString(clear)
			diff.WriteString(trim)
		}
	}
	x, y := e.cursor.GetRealPos(e.starting)
	moveto := fmt.Sprintf("\033[%d;%dH", y, x)
	diff.WriteString(moveto)

	fmt.Print(diff.String())
}
