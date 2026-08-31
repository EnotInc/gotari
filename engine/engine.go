package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.con/enotinc/gotari/enums/cmd"
	"github.con/enotinc/gotari/enums/keys"
	"golang.org/x/term"
)

const termOffset int = 1

type Engine struct {
	// list of games
	Games []*Game

	// map of line indexes to saved hashes
	hash map[int]uint32

	// additional stuff to work with terminal raw state
	fdIn  int
	fdOut int
	old   *term.State
}

type Game interface {
	Render() []*string
	Handle(key rune)
}

func Init() *Engine {
	_fdIn := int(os.Stdin.Fd())
	_fdOut := int(os.Stdout.Fd())

	_old, err := term.MakeRaw(_fdIn)
	if err != nil {
		panic(err)
	}

	return &Engine{
		hash: make(map[int]uint32),

		old:   _old,
		fdIn:  _fdIn,
		fdOut: _fdOut,
	}
}

func (e *Engine) AddGame(g Game) {
	e.Games = append(e.Games, &g)
}

func (e *Engine) exit() {
	var quit strings.Builder
	quit.WriteString(cmd.ClearView)
	quit.WriteString(cmd.ClearHistory)
	quit.WriteString(cmd.MoveToStart)
	quit.WriteString(cmd.ResetTerminal)
	fmt.Print(quit.String())

	term.Restore(e.fdIn, e.old)
}

func (e *Engine) Run() {
	var begin strings.Builder
	begin.WriteString(cmd.SaveTerminal)
	begin.WriteString(cmd.ClearHistory)
	begin.WriteString(cmd.ClearView)
	fmt.Print(begin.String())

	defer e.exit()

	reader := bufio.NewReader(os.Stdin)
	for {
		var diff strings.Builder
		key, _, err := reader.ReadRune()
		if err != nil {
			panic(err)
		}

		// FIXME: quit in menu
		// TODO: figure out how to create menu
		if key == keys.Esc {
			break
		}

		game := *e.Games[0]
		game.Handle(key)
		render := game.Render()

		for index, line := range render {
			curHash := getHash(*line)
			oldHash, ok := e.hash[index]

			if !ok || (ok && curHash != oldHash) {
				pos := fmt.Sprintf("\033[%d;%d;H", index+termOffset, termOffset)
				clear := "\033[0K"

				diff.WriteString(pos)
				diff.WriteString(clear)
				diff.WriteString(*line)
			}
		}
		fmt.Print(diff.String())
	}
}
