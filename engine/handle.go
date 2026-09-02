package engine

import "github.con/enotinc/gotari/engine/cmd"

func (e *Engine) handle(event Event) bool {
	if e.game == nil {
		c := (*e.menu).Handle(event)
		if c != nil {
			switch c {
			case cmd.SelectGame:
				e.clear()
				e.game = (*e.menu).SelectedGame()

			case cmd.QuitGotari:
				return true
			}
		}

	} else {
		c := (*e.game).Handle(event)
		switch c {
		case cmd.CloseGame:
			e.clear()
			e.game = nil
		}
	}

	return false
}
