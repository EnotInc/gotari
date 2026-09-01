package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.con/enotinc/gotari/enums/ascii"
	"github.con/enotinc/gotari/enums/keys"
	"golang.org/x/term"
)

const termOffset int = 1

func Init(menu Menu) *Engine {
	_fdIn := int(os.Stdin.Fd())

	_old, err := term.MakeRaw(_fdIn)
	if err != nil {
		panic(err)
	}

	return &Engine{
		hash: make(map[int]uint32),

		fullscreen: false,

		menu: &menu,

		old:  _old,
		fdIn: _fdIn,
	}
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

		if key == keys.Esc {
			return
		}

		e.handle(key)
		e.render()
	}
}

func (e *Engine) handle(key rune) {
	if e.game == nil {
		game := (*e.menu).Handle(key)
		if game != nil {
			e.game = game
		}

	} else {
		(*e.game).Handle(key)
	}
}
