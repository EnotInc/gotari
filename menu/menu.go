package menu

import (
	"fmt"

	"github.con/enotinc/gotari/engine"
	"github.con/enotinc/gotari/enums/keys"
)

type MainMenu struct {
	games  []*engine.Game
	cursor int
}

func Init() *MainMenu {
	return &MainMenu{
		games:  make([]*engine.Game, 0),
		cursor: 0,
	}
}

func (m *MainMenu) LoadList(games []*engine.Game) {
	for _, game := range games {
		m.games = append(m.games, game)
	}
}

func (m *MainMenu) Render() []*string {
	var render []*string
	for index, game := range m.games {
		c := " "
		if index == m.cursor {
			c = ">"
		}

		name := (*game).Name()
		line := fmt.Sprintf("%s%s", c, name)
		render = append(render, &line)
	}
	return render
}

func (m *MainMenu) Handle(key rune) *engine.Game {
	switch key {
	case 'j':
		if m.cursor < len(m.games)-1 {
			m.cursor += 1
		}
	case 'k':
		if m.cursor > 0 {
			m.cursor -= 1
		}
	case keys.Enter:
		if m.cursor >= 0 && m.cursor <= len(m.games)-1 {
			return m.games[m.cursor]
		}
	}

	return nil
}
