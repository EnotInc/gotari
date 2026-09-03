package engine

import (
	"fmt"
	"os"
	"strings"

	"github.con/enotinc/gotari/engine/cursor"
	"github.con/enotinc/gotari/engine/options"
	"github.con/enotinc/gotari/enums/ascii"
	"golang.org/x/term"
)

const termOffset int = 1

func Init(opt *options.Options, menu Menu) *Engine {
	_fdIn := int(os.Stdin.Fd())

	_old, err := term.MakeRaw(_fdIn)
	if err != nil {
		panic(err)
	}

	c := cursor.NewCursor()
	e := &Engine{
		hash:   make(map[int]uint32),
		menu:   &menu,
		opt:    opt,
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
	if !e.opt.Fullscreen {
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
	if e.opt.Fullscreen {
		begin.WriteString(ascii.SaveTerminal)
		begin.WriteString(ascii.ClearHistory)
		begin.WriteString(ascii.ClearView)
	} else {
		begin.WriteString(ascii.MoveUp)
		begin.WriteString(ascii.SaveCursorPos)
	}

	begin.WriteString(ascii.EnableMouse)
	fmt.Print(begin.String())

	(*e.menu).LoadList(e.Games)
}

func (e *Engine) exit() {
	var quit strings.Builder

	if e.opt.Fullscreen {
		quit.WriteString(ascii.ClearView)
		quit.WriteString(ascii.ClearHistory)
		quit.WriteString(ascii.MoveToStart)
		quit.WriteString(ascii.ResetTerminal)
	} else {
		quit.WriteString(ascii.ResetCursorPos)
		quit.WriteString(ascii.MoveDown)
		quit.WriteString(ascii.ClearAfter)
	}

	quit.WriteString(ascii.DisableMouse)
	quit.WriteString(e.cursor.Reset())
	fmt.Print(quit.String())
	term.Restore(e.fdIn, e.old)

	if r := recover(); r != nil {
		fmt.Print(r)
	}
}

func (e *Engine) Run() {
	e.begin()
	e.render()
	defer e.exit()

	for {
		event := e.getEvent()
		if event == nil {
			continue
		}

		if quit := e.handle(event); quit {
			break
		}
		e.render()
	}
}

func (e *Engine) getEvent() Event {
	buffer := make([]byte, 64)
	n, err := os.Stdin.Read(buffer)
	if err != nil {
		return nil
	}

	input := buffer[:n]
	var event Event

	// -==[ mouse support ]==-
	if e.opt.EnableMouse && n >= 6 && input[0] == '\033' && input[1] == '[' && input[2] == '<' {
		var button, x, y int
		var releaseChar rune

		_, err := fmt.Sscanf(string(input[3:]), "%d;%d;%d%c", &button, &x, &y, &releaseChar)
		if err == nil && releaseChar == 'm' { // on mouse press
			_x := x - e.starting.X - termOffset
			_y := y - e.starting.Y - termOffset
			if _y < 0 || _x < 0 {
				return nil
			}

			event = MouseEvent{X: _x, Y: _y}
		}

	} else

	// -==[ key pressed ]==-
	if n == 1 {
		key := rune(input[0])
		event = KeyEvent{Key: key}

	} else {
		return nil
	}

	return event
}
