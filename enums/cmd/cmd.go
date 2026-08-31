package cmd

const (
	ClearView    = "\033[2J"
	ClearHistory = "\033[3J"
	MoveToStart  = "\033[0H"

	SaveTerminal  = "\033[?1049h"
	ResetTerminal = "\033[?1049l"
)
