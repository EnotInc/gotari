package tictactoe

import "strconv"

type win string

const (
	not  win = "not"
	draw win = "draw"

	diagA win = "\\"
	diagB win = "/"

	row1 win = "r1"
	row2 win = "r2"
	row3 win = "r3"

	col1 win = "c1"
	col2 win = "c2"
	col3 win = "c3"
)

func (t *TTT) isWin() win {
	if t.board[0] == t.board[1] && t.board[1] == t.board[2] && t.board[2] != empty {
		return row1
	}

	if t.board[0] == t.board[4] && t.board[4] == t.board[8] && t.board[4] != empty {
		return diagA
	}

	if t.board[2] == t.board[4] && t.board[4] == t.board[6] && t.board[4] != empty {
		return diagB
	}

	for i := range 3 {
		if t.board[i*3] == t.board[i*3+1] && t.board[i*3+1] == t.board[i*3+2] && t.board[i*3+1] != empty {
			switch i {
			case 0:
				return row1
			case 1:
				return row2
			case 2:
				return row3
			}
		}

		if t.board[i] == t.board[i+1*3] && t.board[i+1*3] == t.board[i+2*3] && t.board[i+1*3] != empty {
			switch i {
			case 0:
				return col1
			case 1:
				return col2
			case 2:
				return col3
			}
		}
	}
	if t.turn >= 9 {
		return draw
	}

	return not
}

func (t *TTT) makeTurn(key rune) (success bool) {
	p, err := strconv.Atoi(string(key))
	if err != nil {
		t.message = "Incorrect input"
		return false
	}
	p -= 1 // fixing offset
	if p < 0 || p > 8 {
		t.message = "Incorrect input"
		return
	}

	var s sqare
	switch t.turn%2 == 0 {
	case true:
		s = cross
	case false:
		s = zero
	}

	if t.board[p] != empty {
		t.message = "this sqare is already taken"
		return false
	}

	t.board[p] = s
	return true
}
