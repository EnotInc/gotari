package menu

import (
	"github.con/enotinc/gotari/engine"
	"github.con/enotinc/gotari/engine/cmd"
	"github.con/enotinc/gotari/engine/cursor"
	"github.con/enotinc/gotari/enums/keys"
)

type MainMenu struct {
	cards  []*card
	cursor int
}

func Init() *MainMenu {
	return &MainMenu{
		cards:  make([]*card, 0),
		cursor: 0,
	}
}

func (m *MainMenu) LoadList(games []*engine.Game) {
	for _, game := range games {
		c := NewCard((*game).Name(), game)
		m.cards = append(m.cards, c)
	}
}

func (m *MainMenu) SelectedGame() *engine.Game {
	return m.cards[m.cursor].game
}

func (m *MainMenu) Render() []*string {
	var render []*string
	title := " List of games:"
	clue := " press <enter> or <space> to select a game. <esc> to quit"
	render = append(render, &title)
	render = append(render, &clue)
	for index, card := range m.cards {
		render = append(render, card.render(index == m.cursor)...)
	}

	return render
}

func (m *MainMenu) Handle(key rune) cmd.CMD {
	switch key {
	case 'j':
		if m.cursor < len(m.cards)-1 {
			m.cursor += 1
		}

	case 'k':
		if m.cursor > 0 {
			m.cursor -= 1
		}

	case keys.Esc:
		return cmd.QuitGotari

	case keys.Enter, keys.Space:
		if m.cursor >= 0 && m.cursor <= len(m.cards)-1 {
			return cmd.SelectGame
		}
	}

	return nil
}

func (m *MainMenu) CursorInfo() (cursor.Kind, cursor.Position) {
	return cursor.Hidden, cursor.Position{}
}
