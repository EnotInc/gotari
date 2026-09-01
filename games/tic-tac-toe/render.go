package tictactoe

import "fmt"

const (
	reset = "\033[0m"

	yellow = "\033[33m"
	green  = "\033[32m"
	gray   = "\033[90m"
	red    = "\033[31m"

	symbol_cross = "x"
	symbol_zero  = "o"
)

func ToStr(s sqare, at int, w win) string {
	if s == empty {
		return fmt.Sprintf(" %s%d%s ", gray, at+1, reset)
	}

	switch w {
	case row1:
		if at == 0 || at == 1 || at == 2 {
			return fmt.Sprintf("%s---%s", yellow, reset)
		}
	case row2:
		if at == 3 || at == 4 || at == 5 {
			return fmt.Sprintf("%s---%s", yellow, reset)
		}
	case row3:
		if at == 6 || at == 7 || at == 8 {
			return fmt.Sprintf("%s---%s", yellow, reset)
		}

	case col1:
		if at == 0 || at == 3 || at == 6 {
			return fmt.Sprintf(" %s|%s ", yellow, reset)
		}
	case col2:
		if at == 1 || at == 4 || at == 7 {
			return fmt.Sprintf(" %s|%s ", yellow, reset)
		}
	case col3:
		if at == 2 || at == 5 || at == 8 {
			return fmt.Sprintf(" %s|%s ", yellow, reset)
		}

	case diagA:
		if at == 0 || at == 4 || at == 8 {
			return fmt.Sprintf(" %s\\%s ", yellow, reset)
		}

	case diagB:
		if at == 2 || at == 4 || at == 6 {
			return fmt.Sprintf(" %s/%s ", yellow, reset)
		}
	}

	switch s {
	case cross:
		return fmt.Sprintf(" %s%s%s ", red, symbol_cross, reset)
	case zero:
		return fmt.Sprintf(" %s%s%s ", green, symbol_zero, reset)
	}
	return ""
}

func (t *TTT) draw() []*string {
	var board []*string
	win := t.isWin()
	for i := range 3 {
		line := fmt.Sprintf("%s|%s|%s",
			ToStr(t.board[3*i], 3*i, win),
			ToStr(t.board[3*i+1], 3*i+1, win),
			ToStr(t.board[3*i+2], 3*i+2, win))

		board = append(board, &line)
		if i != 2 {
			separator := "---+---+---"
			board = append(board, &separator)
		}
	}

	switch win {
	case draw:
		t.message = "draw. Press 'r' to restart"
	case not:
	default:
		switch t.turn%2 != 0 {
		case true:
			t.message = "x win. Press 'r' to restart"
		case false:
			t.message = "o win. Press 'r' to restart"
		}
	}

	board = append(board, &t.message)

	return board
}
