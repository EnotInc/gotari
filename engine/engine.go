package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.con/enotinc/gotari/engine/cmd"
	"github.con/enotinc/gotari/enums/ascii"
	"golang.org/x/term"
)

const termOffset int = 1

func Init(menu Menu) *Engine {
	_fdIn := int(os.Stdin.Fd())

	_old, err := term.MakeRaw(_fdIn)
	if err != nil {
		panic(err)
	}

	c := newCursor()
	e := &Engine{
		// FIXME: move to options, or smth
		fullscreen: false,

		hash:   make(map[int]uint32),
		menu:   &menu,
		old:    _old,
		fdIn:   _fdIn,
		cursor: c,
	}
	e.cursor.changeCursor(Hidden)
	return e
}

func (e *Engine) AddGame(g Game) {
	e.Games = append(e.Games, &g)
}

func (e *Engine) begin() {
	var begin strings.Builder
	if e.fullscreen {
		begin.WriteString(ascii.SaveTerminal)
		begin.WriteString(ascii.ClearHistory)
		begin.WriteString(ascii.ClearView)
	} else {
		begin.WriteString(ascii.MoveUp)
		begin.WriteString(ascii.SaveCursorPos)
	}
	fmt.Print(begin.String())

	(*e.menu).LoadList(e.Games)
}

func (e *Engine) exit() {
	var quit strings.Builder

	if e.fullscreen {
		quit.WriteString(ascii.ClearView)
		quit.WriteString(ascii.ClearHistory)
		quit.WriteString(ascii.MoveToStart)
		quit.WriteString(ascii.ResetTerminal)
	} else {
		quit.WriteString(ascii.ResetCursorPos)
		quit.WriteString(ascii.MoveDown)
		quit.WriteString(ascii.ClearAfter)
	}

	quit.WriteString(_showCursor)
	fmt.Print(quit.String())
	term.Restore(e.fdIn, e.old)
}

func (e *Engine) Run() {
	e.begin()
	e.render()
	defer e.exit()

	reader := bufio.NewReader(os.Stdin)
	for {
		key, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}

		if quit := e.handle(key); quit {
			break
		}
		e.render()
	}
}

// TODO: add mouse and arrows support
func (e *Engine) handle(key rune) bool {
	if e.game == nil {
		c := (*e.menu).Handle(key)
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
		command := (*e.game).Handle(key)
		switch command {
		case cmd.CloseGame:
			e.clear()
			e.game = nil
		}
	}

	return false
}
