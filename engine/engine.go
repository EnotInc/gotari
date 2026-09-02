package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.con/enotinc/gotari/engine/cursor"
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

	c := cursor.NewCursor()
	e := &Engine{
		// FIXME: move to options, or smth
		fullscreen: false,

		hash:   make(map[int]uint32),
		menu:   &menu,
		old:    _old,
		fdIn:   _fdIn,
		cursor: c,
	}
	e.cursor.ChangeCursor(cursor.Hidden)
	e.setStartedPoint()
	return e
}

func (e *Engine) AddGame(g Game) {
	e.Games = append(e.Games, &g)
}

func (e *Engine) setStartedPoint() {
	y := 0
	if !e.fullscreen {
		var row, col int
		fmt.Print("\033[6n")
		_, err := fmt.Fscanf(os.Stdin, "\033[%d;%dR", &row, &col)
		if err != nil {
			panic(err)
		} else {
			y = row - termOffset
		}
	}
	e.starting = cursor.Position{
		X: 0,
		Y: y,
	}
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

	quit.WriteString(e.cursor.Reset())
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
