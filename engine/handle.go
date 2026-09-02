package engine

import "github.con/enotinc/gotari/engine/cmd"

// TODO: add mouse and arrows support
func (e *Engine) handle(key rune) bool {
	if e.game == nil {
		c := (*e.menu).Handle(key)
		if c != nil {
			switch c {
			case cmd.SelectGame:
				e.clear()
				e.game = (*e.menu).SelectedGame()
				kind, pos := (*e.game).CursorInfo()
				e.cursor.ChangeCursor(kind)
				e.cursor.ChangePos(pos)

			case cmd.QuitGotari:
				return true

			default:
				kind, pos := (*e.game).CursorInfo()
				e.cursor.ChangeCursor(kind)
				e.cursor.ChangePos(pos)
			}
		}

	} else {
		command := (*e.game).Handle(key)
		switch command {
		case cmd.CloseGame:
			e.clear()
			e.game = nil

		default:
			kind, pos := (*e.game).CursorInfo()
			e.cursor.ChangeCursor(kind)
			e.cursor.ChangePos(pos)
		}

	}

	return false
}
